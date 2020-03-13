package handler

import (
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("ta04d3ti")

type CustomClaims struct {
	User *userPB.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *userPB.User) (string, error)
}

type TokenService struct{}

// Decode a token string into a token oject
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate the token and return the custom claims
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *userPB.User) (string, error) {
	expire := time.Now().Add(time.Hour * 1).Unix()

	// Create the claims
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "com.ta04.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and return
	return token.SignedString(secretKey)
}
