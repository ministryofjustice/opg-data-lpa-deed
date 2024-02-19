package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/ddb"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/objectstore"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/shared"
	"github.com/ministryofjustice/opg-go-common/logging"
)

type Logger interface {
	Print(...interface{})
}

type Store interface {
	Put(ctx context.Context, data any) error
	Get(ctx context.Context, uid string) (shared.Lpa, error)
}

type StaticLpaStorage interface {
	Save(lpa *shared.Lpa) error
}

type Verifier interface {
	VerifyHeader(events.APIGatewayProxyRequest) (*shared.LpaStoreClaims, error)
}

type Lambda struct {
	store    Store
	staticLpaStorage StaticLpaStorage
	verifier Verifier
	logger   Logger
}

func (l *Lambda) HandleEvent(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := l.verifier.VerifyHeader(event)
	if err != nil {
		l.logger.Print("Unable to verify JWT from header")
		return shared.ProblemUnauthorisedRequest.Respond()
	}

	l.logger.Print("Successfully parsed JWT from event header")

	var input shared.LpaInit
	uid := event.PathParameters["uid"]

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

	err = json.Unmarshal([]byte(event.Body), &input)
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

	// save to static storage as JSON
	err = l.staticLpaStorage.Save(&data)

	if err != nil {
		l.logger.Print(err)
		return shared.ProblemInternalServerError.Respond()
	}

	// respond
	response.StatusCode = 201
	response.Body = `{}`

	return response, nil
}

func main() {
	l := &Lambda{
		store: ddb.New(
			os.Getenv("AWS_DYNAMODB_ENDPOINT"),
			os.Getenv("DDB_TABLE_NAME_DEEDS"),
			os.Getenv("DDB_TABLE_NAME_CHANGES"),
		),
		staticLpaStorage: objectstore.NewStaticLpaStorage(
			objectstore.NewS3Client(
				os.Getenv("S3_BUCKET_NAME_ORIGINAL"),
				os.Getenv("AWS_S3_ENDPOINT"),
			),
		),
		verifier: shared.NewJWTVerifier(),
		logger: logging.New(os.Stdout, "opg-data-lpa-store"),
	}

	lambda.Start(l.HandleEvent)
}
