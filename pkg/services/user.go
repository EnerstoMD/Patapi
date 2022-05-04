package service

import (
	"errors"
	"fmt"
	"lupus/patapi/pkg/auth"
	"lupus/patapi/pkg/model"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(c *gin.Context, u model.User) error
	Login(c *gin.Context, u model.User) (string, error)
	VerifyUserExists(c *gin.Context, u model.User) error
	Logout(c *gin.Context) error
	GetUserInfo(c *gin.Context) (model.User, error)
}

type UserDb interface {
	CreateUser(c *gin.Context, u model.User) error
	GetUserByEmail(c *gin.Context, email string) (model.User, error)
	VerifyUserExists(c *gin.Context, u model.User) error
	GetUserById(c *gin.Context, id string) (model.User, error)
}
type userService struct {
	d UserDb
	a auth.AuthService
}

func NewUserService(d UserDb, a auth.AuthService) UserService {
	return &userService{d, a}
}

func (s *userService) CreateUser(c *gin.Context, u model.User) error {
	err := u.ValidateUser(u)
	if err != nil {
		return err
	}
	*u.Password, err = u.EncryptPassword(*u.Password)
	if err != nil {
		return err
	}
	err = s.VerifyUserExists(c, u)
	if err != nil {
		return err
	}
	return s.d.CreateUser(c, u)
}

func (s *userService) Login(c *gin.Context, u model.User) (string, error) {
	searchedUser, err := s.d.GetUserByEmail(c, *u.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*searchedUser.Password), []byte(*u.Password)); err != nil {
		return "", err
	}

	jwtWrapper := model.JwtWrapper{
		SecretKey:       "secret",
		Issuer:          "lupus",
		ExpirationHours: 24,
	}

	return s.a.GenerateToken(c, *searchedUser.Id, *searchedUser.Email, jwtWrapper)
}

func (s *userService) VerifyUserExists(c *gin.Context, u model.User) error {
	return s.d.VerifyUserExists(c, u)
}

func (s *userService) Logout(c *gin.Context) error {
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	if token == "" {
		return errors.New("no token, cannot logout")
	}
	return s.a.DeleteToken(c, token)
}

func (s *userService) GetUserInfo(c *gin.Context) (model.User, error) {
	user, err := s.d.GetUserById(c, fmt.Sprintf("%v", c.Keys["userId"]))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
