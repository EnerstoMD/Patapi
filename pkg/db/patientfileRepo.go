package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

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

func (repo *DbSources) RegisterPatientDisease(c *gin.Context, p model.PatientDisease) error {
	dQueryArgs, err := utils.PrepareSQLInsertStatement(c, p.Disease)
	if err != nil {
		return err
	}

	patDqueryArgs, err := utils.ReadStructToBeInserted(c, p)
	if err != nil {
		return err
	}
	var diseaseId int64
	var dId string

	if p.Disease.Id != nil {
		foundDisease, err := repo.FindDiseaseById(c, *p.Disease.Id)
		if err != nil {
			return err
		}
		if foundDisease.Id == nil {
			row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
			err = row.Scan(&diseaseId)
			if err != nil {
				return err
			}
			dId = strconv.Itoa(int(diseaseId))
		} else {
			dId = *foundDisease.Id
		}
	} else {
		row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
		err = row.Scan(&diseaseId)
		if err != nil {
			return err
		}
		dId = strconv.Itoa(int(diseaseId))
	}

	columnArgs := utils.InsertInsideQueryString(patDqueryArgs[1], "disease_id")
	valuesArgs := utils.InsertInsideQueryString(patDqueryArgs[3], dId)
	query := `INSERT INTO ` + patDqueryArgs[0] + ` ` + columnArgs + ` VALUES ` + valuesArgs
	log.Print(query)
	tx, err := repo.dbConn.Begin()
	if err != nil {
		log.Println("db Begin() miss")
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	if _, err = tx.Exec(query); err != nil {
		return err
	}
	return err
}

func (repo *DbSources) FindDiseaseById(c *gin.Context, id string) (model.Disease, error) {
	var d model.Disease
	err := repo.dbConn.Get(&d, `SELECT * FROM disease WHERE id = $1`, id)
	return d, err
}

func (repo *DbSources) GetPatientDiseases(c *gin.Context, patientId string) ([]model.PatientDisease, error) {
	var diseases []model.PatientDisease
	rows, err := repo.dbConn.Queryx(`SELECT id,patient_id,start_date,end_date,comment,in_progress,added_by,disease_id FROM patient_disease WHERE patient_id = $1`, patientId)
	//err := repo.dbConn.Select(&diseases, `SELECT * FROM patient_disease WHERE patient_id = $1`, patientId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.PatientDisease
		var diseaseId string

		err = rows.Scan(&d.Id, &d.PatientId, &d.StartDate, &d.EndDate, &d.Comment, &d.InProgress, &d.AddedBy, &diseaseId)
		if err != nil {
			return nil, err
		}

		d.Disease, err = repo.FindDiseaseById(c, diseaseId)
		if err != nil {
			return nil, err
		}
		diseases = append(diseases, d)
	}
	return diseases, err
}

func (repo *DbSources) DeletePatientDisease(c *gin.Context, patientId, id string) error {
	_, err := repo.dbConn.Exec(`DELETE FROM patient_disease WHERE patient_id = $1 AND id = $2`, patientId, id)
	return err
}

func (repo *DbSources) UpdatePatientDisease(c *gin.Context, p model.PatientDisease) error {
	updDquery, err := utils.PrepareSQLUpdateStatement(p.Disease, *p.Disease.Id)
	if err != nil {
		return err
	}
	err = repo.execQuery(updDquery)
	if err != nil {
		return err
	}
	updPatDquery, err := utils.PrepareSQLUpdateStatement(p, *p.Id)
	if err != nil {
		return err
	}
	return repo.execQuery(updPatDquery)
}
