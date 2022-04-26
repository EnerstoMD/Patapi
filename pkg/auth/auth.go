package auth

import (
	"errors"
	"lupus/patapi/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

type AuthService interface {
	GenerateToken(c *gin.Context, userId, email string, j model.JwtWrapper) (string, error)
	ValidateToken(c *gin.Context, tokenString string, j model.JwtWrapper) (*model.JwtClaims, error)
}

type TokenDb interface {
	SetRefreshToken(c *gin.Context, userID, tokenID string, expiresIn time.Duration) error
	ValidateToken(c *gin.Context, userID, previoustokenID string) error
}

type authService struct {
	t TokenDb
}

func NewAuthService(t TokenDb) AuthService {
	return &authService{t}
}

func (auth *authService) GenerateToken(c *gin.Context, userId string, email string, j model.JwtWrapper) (string, error) {
	claims := &model.JwtClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	err2 := auth.t.SetRefreshToken(c, email, tokenStr, time.Hour*time.Duration(j.ExpirationHours))
	if err != nil {
		return "", err2
	}
	return tokenStr, nil
}

func (auth *authService) ValidateToken(c *gin.Context, tokenString string, j model.JwtWrapper) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	if claims.Issuer != j.Issuer {
		return nil, errors.New("invalid issuer")
	}
	err2 := auth.t.ValidateToken(c, claims.Email, tokenString)
	if err2 != nil {
		return nil, err2
	}

	return claims, err
}
