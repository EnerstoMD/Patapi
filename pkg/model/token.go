package model

import "github.com/golang-jwt/jwt"

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	Email string `json:"email"`
	Roles []int  `json:"roles"`
	jwt.StandardClaims
}
