package handler

import (
	"fmt"
	"lupus/patapi/pkg/model"
	user "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	AdminUpdatePassword(ctx *gin.Context)
}

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (userHandler *userHandler) Register(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read user", "error": err.Error()})
		return
	}

	err := userHandler.userService.CreateUser(c, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't insert user into db", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "User registered"})
}

func (userHandler *userHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read user", "error": err.Error()})
		return
	}

	token, err := userHandler.userService.Login(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't login user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "token": token})
}

func (userHandler *userHandler) Logout(c *gin.Context) {
	err := userHandler.userService.Logout(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't logout user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "User logged out"})
}

func (userHandler *userHandler) GetUserInfo(c *gin.Context) {
	user, err := userHandler.userService.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't find user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userHandler *userHandler) GetUsers(c *gin.Context) {
	users, err := userHandler.userService.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't find users", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userHandler *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := userHandler.userService.DeleteUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't delete user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "User deleted"})
}

func (userHandler *userHandler) AdminUpdatePassword(c *gin.Context) {
	pwdModel := model.PasswordUpdate{}
	var userId, pwd string

	if err := c.ShouldBindJSON(&pwdModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read password in body", "error": err.Error()})
		return
	}
	userId = c.Param("id")
	pwd = *pwdModel.Password
	if pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read password in body", "error": "password is empty"})
		return
	}

	err := userHandler.userService.UpdatePassword(c, pwd, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't update password", "error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "Password updated"})
}

func (userHandler *userHandler) UpdatePassword(c *gin.Context) {
	pwdModel := model.PasswordUpdate{}
	var userId, pwd string

	if err := c.ShouldBindJSON(&pwdModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read password in body", "error": err.Error()})
		return
	}

	userId = fmt.Sprintf("%v", c.Keys["userId"])
	pwd = *pwdModel.Password
	if pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read password in body", "error": "password is empty"})
		return
	}

	err := userHandler.userService.UpdatePassword(c, pwd, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "can't update password", "error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "Password updated"})
}
