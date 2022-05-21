package db

import (
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) CreatePatientComment(ctx *gin.Context, comment model.PatientComment) error {
	queryArgs, err := utils.ReadStructToBeInserted(ctx, comment)
	if err != nil {
		return err
	}
	_, err = repo.dbConn.NamedExec(`INSERT INTO `+queryArgs[0]+` `+queryArgs[1]+` VALUES `+queryArgs[2], comment)
	return err
}
