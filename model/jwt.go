package model

import "github.com/dgrijalva/jwt-go"

// Claims is the struct of claims needed for JWT
type Claims struct {
	jwt.StandardClaims
	UserID           int32
	Username         string
	UserEmailAddress string
	UserRole         string
}