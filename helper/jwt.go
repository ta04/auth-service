package helper

import (
	"github.com/ta04/auth-service/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ta04/auth-service/model"
	proto "github.com/ta04/user-service/model/proto"
)

// Encode will encode the claims into a JWT
func Encode(user *proto.User) (string, error) {
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(time.Hour * 1).Unix()

	issuer := config.MicroServiceName()

	claims := model.Claims{
		ID:           user.Id,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Role:         user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.SecretKey()

	return token.SignedString([]byte(secretKey))
}
