package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/ddb"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/event"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/objectstore"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/shared"
	"github.com/ministryofjustice/opg-go-common/logging"
)

type EventClient interface {
	SendLpaUpdated(ctx context.Context, event event.LpaUpdated) error
}

type Logger interface {
	Print(...interface{})
}

type Store interface {
	Put(ctx context.Context, data any) error
	Get(ctx context.Context, uid string) (shared.Lpa, error)
}

type S3Client interface {
	Put(objectKey string, obj any) (*s3.PutObjectOutput, error)
	Get(objectKey string) (*s3.GetObjectOutput, error)
}

type Verifier interface {
	VerifyHeader(events.APIGatewayProxyRequest) (*shared.LpaStoreClaims, error)
}

type Lambda struct {
	eventClient EventClient
	s3client    S3Client
	store       Store
	verifier    Verifier
	logger      Logger
}

func (l *Lambda) HandleEvent(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := l.verifier.VerifyHeader(req)
	if err != nil {
		l.logger.Print("Unable to verify JWT from header")
		return shared.ProblemUnauthorisedRequest.Respond()
	}

	l.logger.Print("Successfully parsed JWT from event header")

	var input shared.LpaInit
	uid := req.PathParameters["uid"]

	response := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "{\"code\":\"INTERNAL_SERVER_ERROR\",\"detail\":\"Internal server error\"}",
	}

	// check for existing Lpa
	var existingLpa shared.Lpa
	existingLpa, err = l.store.Get(ctx, uid)
	if err != nil {
		l.logger.Print(err)
		return shared.ProblemInternalServerError.Respond()
	}

	if existingLpa.Uid == uid {
		problem := shared.ProblemInvalidRequest
		problem.Detail = "LPA with UID already exists"
		return problem.Respond()
	}

	err = json.Unmarshal([]byte(req.Body), &input)
	if err != nil {
		l.logger.Print(err)
		return shared.ProblemInternalServerError.Respond()
	}

	// validation
	errs := Validate(input)
	if len(errs) > 0 {
		problem := shared.ProblemInvalidRequest
		problem.Errors = errs

		return problem.Respond()
	}

	data := shared.Lpa{LpaInit: input}
	data.Uid = uid
	data.Status = shared.LpaStatusProcessing
	data.UpdatedAt = time.Now()

	// save
	err = l.store.Put(ctx, data)
	if err != nil {
		l.logger.Print(err)
		return shared.ProblemInternalServerError.Respond()
	}

	// save a copy of the original to permanent storage,
	// but only if the key doesn't already exist
	objectKey := fmt.Sprintf("%s/donor-executed-lpa.json", data.Uid)
	_, err = l.s3client.Get(objectKey)
	if err == nil {
		// 200 => bad (object already exists)
		err = fmt.Errorf("Could not save donor executed LPA as key %s already exists", objectKey)
		l.logger.Print(err)
		return shared.ProblemInvalidRequest.Respond()
	}

	// 404 => good (object should not already exist)
	var nsk *types.NoSuchKey
	if errors.As(err, &nsk) {
		_, err = l.s3client.Put(objectKey, data)
    }

	if err != nil {
		l.logger.Print(err)
		return shared.ProblemInternalServerError.Respond()
	}

	// send lpa-updated event
	err = l.eventClient.SendLpaUpdated(ctx, event.LpaUpdated{
		Uid: uid,
		ChangeType: "CREATED",
	})

	if err != nil {
		l.logger.Print(err)
	}

	// respond
	response.StatusCode = 201
	response.Body = `{}`

	return response, nil
}

func main() {
	logger := logging.New(os.Stdout, "opg-data-lpa-store")
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
	  logger.Print("Failed to load configuration:", err)
	}

	l := &Lambda{
		eventClient: event.NewClient(awsConfig, os.Getenv("EVENT_BUS_NAME")),
		store:    ddb.New(
			os.Getenv("AWS_DYNAMODB_ENDPOINT"),
			os.Getenv("DDB_TABLE_NAME_DEEDS"),
			os.Getenv("DDB_TABLE_NAME_CHANGES"),
		),
		s3client: objectstore.NewS3Client(
			os.Getenv("S3_BUCKET_NAME_ORIGINAL"),
			os.Getenv("AWS_S3_ENDPOINT"),
		),
		verifier: shared.NewJWTVerifier(),
		logger:   logger,
	}

	lambda.Start(l.HandleEvent)
}
