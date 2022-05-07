package model

import "github.com/golang-jwt/jwt"

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles []int  `json:"roles"`
	jwt.StandardClaims
}
