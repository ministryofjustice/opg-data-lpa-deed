package shared

import (
	"os"
	"testing"
	"time"

    "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var secretKey = []byte("secret")

var verifier = JWTVerifier{
	secretKey: secretKey,
}

func createToken(claims jwt.MapClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secretKey)

    if err != nil {
    	return "", err
    }

 	return tokenString, nil
}

func TestVerifyEmptyJwt(t *testing.T) {
	err := verifier.VerifyToken("")
	assert.NotNil(t, err)
}

func TestVerifyExpInPast(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * -24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "opg.poas.makeregister",
        "sub": "M-3467-89QW-ERTY",
    })

	err := verifier.VerifyToken(token)

	assert.NotNil(t, err)
	if err != nil {
		assert.Containsf(t, err.Error(), "token is expired", "")
	}
}

func TestVerifyIatInFuture(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * 24).Unix(),
        "iss": "opg.poas.sirius",
        "sub": "someone@someplace.somewhere.com",
    })

	err := verifier.VerifyToken(token)

	assert.NotNil(t, err)
	if err != nil {
		assert.Containsf(t, err.Error(), "IssuedAt must not be in the future", "")
	}
}

func TestVerifyIssuer(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "daadsdaadsadsads",
        "sub": "someone@someplace.somewhere.com",
    })

	err := verifier.VerifyToken(token)

	assert.NotNil(t, err)
	if err != nil {
		assert.Containsf(t, err.Error(), "Invalid Issuer", "")
	}
}

func TestVerifyBadEmailForSiriusIssuer(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "opg.poas.sirius",
        "sub": "",
    })

	err := verifier.VerifyToken(token)

	assert.NotNil(t, err)
	if err != nil {
		assert.Containsf(t, err.Error(), "Subject is not a valid email", "")
	}
}

func TestVerifyBadUIDForMRLPAIssuer(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "opg.poas.makeregister",
        "sub": "",
    })

	err := verifier.VerifyToken(token)

	assert.NotNil(t, err)
	if err != nil {
		assert.Containsf(t, err.Error(), "Subject is not a valid UID", "")
	}
}

func TestVerifyGoodJwt(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "opg.poas.sirius",
        "sub": "someone@someplace.somewhere.com",
    })

    err := verifier.VerifyToken(token)
	assert.Nil(t, err)
}

func TestNewJWTVerifier(t *testing.T) {
	token, _ := createToken(jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 24).Unix(),
        "iat": time.Now().Add(time.Hour * -24).Unix(),
        "iss": "opg.poas.sirius",
        "sub": "someone@someplace.somewhere.com",
    })

    os.Setenv("JWT_SECRET_KEY", string(secretKey))
    newVerifier := NewJWTVerifier()
    os.Unsetenv("JWT_SECRET_KEY")

    err := newVerifier.VerifyToken(token)
	assert.Nil(t, err)
}

