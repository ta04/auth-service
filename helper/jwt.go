package helper

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ta04/auth-service/model"
	proto "github.com/ta04/user-service/model/proto"
)

// Encode will encode the claims into a JWT
func Encode(user *proto.User) (string, error) {
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 1).Unix()

	issuer, exist := os.LookupEnv("MICRO_SERVICE_NAME")
	if !exist {
		return "", errors.New("MICRO_WEB_NAME is not exists in .env file")
	}

	claims := model.Claims{
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
