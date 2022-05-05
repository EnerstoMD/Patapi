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

func (repo *DbSources) GetUserByEmail(c *gin.Context, email string) (model.User, error) {
	query := `SELECT * FROM public.user WHERE email='` + email + `'`
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
	return model.User{}, errors.New("user email is incorrect")
}

func (repo *DbSources) GetUserById(c *gin.Context, id string) (model.User, error) {
	query := `SELECT id,name,email FROM public.user WHERE id='` + id + `'`
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
	return model.User{}, errors.New("unknown user id")
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
	return errors.New("user email already exists")
}

func (repo *DbSources) GetUsers(c *gin.Context) ([]model.User, error) {
	query := `SELECT id,name,email FROM public.user`
	rows, err := repo.dbConn.Queryx(query)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *DbSources) DeleteUser(c *gin.Context, id string) error {
	query := `DELETE FROM public.user WHERE id='` + id + `'`
	return repo.execQuery(query)
}

func (repo *DbSources) UpdateUser(c *gin.Context, u model.User) error {
	query, err := utils.PrepareSQLUpdateStatement(u, *u.Id)
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}
