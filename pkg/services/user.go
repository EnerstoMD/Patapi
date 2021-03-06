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
	GetUsers(c *gin.Context) ([]model.User, error)
	DeleteUser(c *gin.Context, id string) error
	UpdatePassword(c *gin.Context, password, userId string) error
	UpdateUserInfo(c *gin.Context, u model.User) error
	GetUserRoles(c *gin.Context, userId string) ([]int, error)
}

type UserDb interface {
	CreateUser(c *gin.Context, u model.User) error
	GetUserByEmail(c *gin.Context, email string) (model.User, error)
	VerifyUserExists(c *gin.Context, u model.User) error
	GetUserById(c *gin.Context, id string) (model.User, error)
	GetUsers(c *gin.Context) ([]model.User, error)
	DeleteUser(c *gin.Context, id string) error
	UpdateUser(c *gin.Context, u model.User) error
	GetUserRoles(c *gin.Context, id string) ([]int, error)
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

func (s *userService) UpdatePassword(c *gin.Context, password, userId string) error {
	user, err := s.d.GetUserById(c, userId)
	if err != nil {
		return err
	}
	encryptedPassword, err := user.EncryptPassword(password)
	if err != nil {
		return err
	}
	user.Password = &encryptedPassword
	return s.d.UpdateUser(c, user)
}

func (s *userService) Login(c *gin.Context, u model.User) (string, error) {
	searchedUser, err := s.d.GetUserByEmail(c, *u.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*searchedUser.Password), []byte(*u.Password)); err != nil {
		return "", err
	}
	roles, err := s.d.GetUserRoles(c, *searchedUser.Id)
	if err != nil {
		return "", err
	}

	jwtWrapper := model.JwtWrapper{
		SecretKey:       "secret",
		Issuer:          "lupus",
		ExpirationHours: 24,
	}

	return s.a.GenerateToken(c, *searchedUser.Name, *searchedUser.Id, *searchedUser.Email, roles, jwtWrapper)
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

func (s *userService) GetUsers(c *gin.Context) ([]model.User, error) {
	return s.d.GetUsers(c)
}

func (s *userService) DeleteUser(c *gin.Context, id string) error {
	//can't delete user 1
	if id == c.Keys["userId"] || id == "1" {
		return errors.New("cannot delete yourself")
	}
	return s.d.DeleteUser(c, id)
}

func (s *userService) UpdateUserInfo(c *gin.Context, u model.User) error {
	return s.d.UpdateUser(c, u)
}

func (s *userService) GetUserRoles(c *gin.Context, userId string) ([]int, error) {
	return s.d.GetUserRoles(c, userId)
}
