package model

import "github.com/dgrijalva/jwt-go"

// Claims is the struct of claims needed for JWT
type Claims struct {
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
	Role         string `json:"role"`
	jwt.StandardClaims
}
