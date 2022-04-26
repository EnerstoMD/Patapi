package db

import (
	"errors"
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) CreateUser(c *gin.Context, u model.User) error {
	query, err := utils.PrepareSQLInsertStatement(c, u)
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}

func (repo *DbSources) GetUserByEmail(c *gin.Context, u model.User) (model.User, error) {
	query := `SELECT * FROM public.user WHERE email='` + *u.Email + `'`
	rows, err := repo.dbConn.Queryx(query)
	var user model.User
	if err != nil {
		return model.User{}, err
	}
	if rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	}
	return model.User{}, errors.New("User email or password is incorrect")
}

func (repo *DbSources) VerifyUserExists(c *gin.Context, u model.User) error {
	query := `SELECT email FROM public.user WHERE email='` + *u.Email + `'`
	log.Println(query)
	rows, err := repo.dbConn.Queryx(query)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return nil
	}
	return errors.New("User email already exists")
}
