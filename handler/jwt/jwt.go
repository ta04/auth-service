package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	userPB "github.com/ta04/user-service/proto"
)

// CustomClaims is the claims needed for JWT
type CustomClaims struct {
	jwt.StandardClaims
	UserID           int32
	Username         string
	UserEmailAddress string
	UserRole         string
}

// JWT is the implementor of TokenHandler interface
type JWT struct{}

// Encode will encode the claims into a JWT
func (srv *JWT) Encode(user *userPB.User) (string, error) {
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 1).Unix()

	issuer, exist := os.LookupEnv("MICRO_SERVICE_NAME")
	if !exist {
		return "", errors.New("MICRO_WEB_NAME is not exists in .env file")
	}

	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			Issuer:    issuer,
		},
		UserID:           user.Id,
		Username:         user.Username,
		UserEmailAddress: user.EmailAddress,
		UserRole:         user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey, exist := os.LookupEnv("SECRET_KEY")
	if !exist {
		return "", errors.New("SECRET_KEY is not exists in .env file")
	}
	return token.SignedString([]byte(secretKey))
}
