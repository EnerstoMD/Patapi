package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       *string `json:"id" db:"id"`
	Name     *string `json:"name" db:"name"`
	Email    *string `json:"email" db:"email"`
	Password string  `json:"password" db:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (user *User) ValidateUser(u User) error {
	if u.Email == nil {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (user *User) EncryptPassword(password string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt), err
}
