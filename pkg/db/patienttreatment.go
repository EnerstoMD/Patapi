package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) RegisterPatientTreatment(c *gin.Context, p model.PatientTreatment) error {
	dQueryArgs, err := utils.PrepareSQLInsertStatement(c, p.Treatment)
	if err != nil {
		return err
	}

	patDqueryArgs, err := utils.ReadStructToBeInserted(c, p)
	if err != nil {
		return err
	}
	var treatmentId int64
	var dId string

	if p.Treatment.Id != nil {
		foundTreatment, err := repo.FindTreatmentById(c, *p.Treatment.Id)
		if err != nil {
			return err
		}
		if foundTreatment.Id == nil {
			row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
			err = row.Scan(&treatmentId)
			if err != nil {
				return err
			}
			dId = strconv.Itoa(int(treatmentId))
		} else {
			dId = *foundTreatment.Id
		}
	} else {
		row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
		err = row.Scan(&treatmentId)
		if err != nil {
			return err
		}
		dId = strconv.Itoa(int(treatmentId))
	}

	columnArgs := utils.InsertInsideQueryString(patDqueryArgs[1], "treatment_id")
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

func (repo *DbSources) FindTreatmentById(c *gin.Context, id string) (model.Treatment, error) {
	var d model.Treatment
	err := repo.dbConn.Get(&d, `SELECT * FROM treatment WHERE id = $1`, id)
	return d, err
}

func (repo *DbSources) GetPatientTreatments(c *gin.Context, patientId string) ([]model.PatientTreatment, error) {
	var treatments []model.PatientTreatment
	rows, err := repo.dbConn.Queryx(`SELECT id,patient_id,start_date,end_date,comment,in_progress,added_by,treatment_id FROM patient_treatment WHERE patient_id = $1`, patientId)
	//err := repo.dbConn.Select(&treatments, `SELECT * FROM patient_treatment WHERE patient_id = $1`, patientId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.PatientTreatment
		var treatmentId string

		err = rows.Scan(&d.Id, &d.PatientId, &d.StartDate, &d.EndDate, &d.Comment, &d.InProgress, &d.AddedBy, &treatmentId)
		if err != nil {
			return nil, err
		}

		d.Treatment, err = repo.FindTreatmentById(c, treatmentId)
		if err != nil {
			return nil, err
		}
		treatments = append(treatments, d)
	}
	return treatments, err
}

func (repo *DbSources) DeletePatientTreatment(c *gin.Context, patientId, id string) error {
	_, err := repo.dbConn.Exec(`DELETE FROM patient_treatment WHERE patient_id = $1 AND id = $2`, patientId, id)
	return err
}

func (repo *DbSources) UpdatePatientTreatment(c *gin.Context, p model.PatientTreatment) error {
	updDquery, err := utils.PrepareSQLUpdateStatement(p.Treatment, *p.Treatment.Id)
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
