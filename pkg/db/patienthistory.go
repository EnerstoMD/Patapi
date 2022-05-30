package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) RegisterPatientHistory(c *gin.Context, p model.PatientHistory) error {
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

func (repo *DbSources) FindHistoryById(c *gin.Context, id string) (model.PatientHistory, error) {
	var d model.PatientHistory
	err := repo.dbConn.Get(&d, `SELECT * FROM patient_history WHERE id = $1`, id)
	return d, err
}

func (repo *DbSources) GetPatientHistory(c *gin.Context, patientId string) ([]model.PatientHistory, error) {
	var historys []model.PatientHistory
	rows, err := repo.dbConn.Queryx(`SELECT id,patient_id,disease_id,family_connection,comment FROM patient_history WHERE patient_id = $1`, patientId)
	//err := repo.dbConn.Select(&historys, `SELECT * FROM patient_history WHERE patient_id = $1`, patientId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.PatientHistory
		var diseaseId string

		err = rows.Scan(&d.Id, &d.PatientId, &diseaseId, &d.FamilyConnection, &d.Comment)
		if err != nil {
			return nil, err
		}

		d.Disease, err = repo.FindDiseaseById(c, diseaseId)
		if err != nil {
			return nil, err
		}
		historys = append(historys, d)
	}
	return historys, err
}

func (repo *DbSources) DeletePatientHistory(c *gin.Context, patientId, id string) error {
	_, err := repo.dbConn.Exec(`DELETE FROM patient_history WHERE patient_id = $1 AND id = $2`, patientId, id)
	return err
}

func (repo *DbSources) UpdatePatientHistory(c *gin.Context, p model.PatientHistory) error {
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
