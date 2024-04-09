package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/session"
	v4old "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ./api-test/tester UID -> generate a UID
// ./api-test/tester JWT -> generate a JWT
// JWT_SECRET_KEY=secret ./api-test/tester -expectedStatus=200 REQUEST <METHOD> <URL> <REQUEST BODY>
//
//	-> make a test request with a JWT generated using secret "secret" and expected status 200
//
// note that the jwtSecret sends a boilerplate JWT for now with valid iat, exp, iss and sub fields
func main() {
	ctx := context.Background()
	expectedStatusCode := flag.Int("expectedStatus", 200, "Expected response status code")
	flag.Parse()
	args := flag.Args()

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	// early exit if we're just generating a UID or JWT
	if args[0] == "UID" {
		fmt.Print("M-" + strings.ToUpper(uuid.NewString()[9:23]))
		os.Exit(0)
	}

	if args[0] == "JWT" {
		fmt.Print(makeJwt([]byte(jwtSecret)))
		os.Exit(0)
	}

	if args[0] != "REQUEST" {
		panic("Unrecognised command")
	}

	method := args[1]
	url := args[2]
	var body io.Reader
	if method != http.MethodGet {
		body = strings.NewReader(args[3])
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	oldreq, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	if body != nil {
		req.Header.Add("Content-type", "application/json")
		oldreq.Header.Add("Content-type", "application/json")
	}

	if jwtSecret != "" {
		tokenString := makeJwt([]byte(jwtSecret))

		req.Header.Add("X-Jwt-Authorization", fmt.Sprintf("Bearer %s", tokenString))
		oldreq.Header.Add("X-Jwt-Authorization", fmt.Sprintf("Bearer %s", tokenString))
	}

	if !strings.HasPrefix(url, "http://localhost") {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic(err)
		}

		signer := v4.NewSigner()

		credentials, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			panic(err)
		}

		hash := sha256.New()
		if body != nil {
			if _, err := io.Copy(hash, body); err != nil {
				panic(err)
			}
		}

		encodedBody := hex.EncodeToString(hash.Sum(nil))

		if err := signer.SignHTTP(ctx, credentials, req, encodedBody, "execute-api", cfg.Region, time.Now()); err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		log.Println("-- SIGNED REQUEST --")
		_ = req.Clone(ctx).Write(&buf)
		log.Println(buf.String())

		// OLD style
		oldsess := session.Must(session.NewSession())
		oldsigner := v4old.NewSigner(oldsess.Config.Credentials)

		_, err = oldsigner.Sign(oldreq, body, "execute-api", "eu-west-1", time.Now())

		buf.Reset()
		log.Println("-- OLD SIGNED REQUEST --")
		_ = oldreq.Clone(ctx).Write(&buf)
		log.Println(buf.String())
	}

	client := http.Client{}
	resp, err := client.Do(oldreq)
	if err != nil {
		panic(err)
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, resp.Body)

	log.Printf("*******************")

	if resp.StatusCode != *expectedStatusCode {
		log.Printf("! TEST FAILED - %s to %s", method, url)
		log.Printf("invalid status code %d; expected: %d", resp.StatusCode, *expectedStatusCode)
		log.Printf("error response: %s", buf.String())

		os.Exit(1)
	} else {
		log.Print(resp.Header)
		log.Printf("Test passed - %s to %s - %d: %s", method, url, resp.StatusCode, buf.String())
	}
}

func makeJwt(secretKey []byte) string {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Add(time.Hour * -24).Unix(),
		"iss": "opg.poas.sirius",
		"sub": "someone@someplace.somewhere.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		panic(err)
	}

	return tokenString
}
