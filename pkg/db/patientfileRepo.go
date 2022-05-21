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

func (repo *DbSources) GetPatientComments(ctx *gin.Context, patientId string) ([]model.PatientComment, error) {
	var comments []model.PatientComment
	err := repo.dbConn.Select(&comments, `SELECT * FROM patient_comment WHERE patient_id = $1`, patientId)
	return comments, err
}

func (repo *DbSources) DeletePatientComment(ctx *gin.Context, patientId, commentId string) error {
	_, err := repo.dbConn.Exec(`DELETE FROM patient_comment WHERE patient_id = $1 AND id = $2`, patientId, commentId)
	return err
}
