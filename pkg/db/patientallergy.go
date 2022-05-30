package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) RegisterPatientAllergy(c *gin.Context, p model.PatientAllergy) error {
	dQueryArgs, err := utils.PrepareSQLInsertStatement(c, p.Allergy)
	if err != nil {
		return err
	}

	patDqueryArgs, err := utils.ReadStructToBeInserted(c, p)
	if err != nil {
		return err
	}
	var allergyId int64
	var dId string

	if p.Allergy.Id != nil {
		foundAllergy, err := repo.FindAllergyById(c, *p.Allergy.Id)
		if err != nil {
			return err
		}
		if foundAllergy.Id == nil {
			row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
			err = row.Scan(&allergyId)
			if err != nil {
				return err
			}
			dId = strconv.Itoa(int(allergyId))
		} else {
			dId = *foundAllergy.Id
		}
	} else {
		row := repo.dbConn.QueryRow(dQueryArgs + ` RETURNING id`)
		err = row.Scan(&allergyId)
		if err != nil {
			return err
		}
		dId = strconv.Itoa(int(allergyId))
	}

	columnArgs := utils.InsertInsideQueryString(patDqueryArgs[1], "allergy_id")
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

func (repo *DbSources) FindAllergyById(c *gin.Context, id string) (model.Allergy, error) {
	var d model.Allergy
	err := repo.dbConn.Get(&d, `SELECT * FROM allergy WHERE id = $1`, id)
	return d, err
}

func (repo *DbSources) GetPatientAllergies(c *gin.Context, patientId string) ([]model.PatientAllergy, error) {
	var allergys []model.PatientAllergy
	rows, err := repo.dbConn.Queryx(`SELECT id,patient_id,start_date,end_date,comment,in_progress,added_by,allergy_id FROM patient_allergy WHERE patient_id = $1`, patientId)
	//err := repo.dbConn.Select(&allergys, `SELECT * FROM patient_allergy WHERE patient_id = $1`, patientId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.PatientAllergy
		var allergyId string

		err = rows.Scan(&d.Id, &d.PatientId, &d.StartDate, &d.EndDate, &d.Comment, &d.InProgress, &d.AddedBy, &allergyId)
		if err != nil {
			return nil, err
		}

		d.Allergy, err = repo.FindAllergyById(c, allergyId)
		if err != nil {
			return nil, err
		}
		allergys = append(allergys, d)
	}
	return allergys, err
}

func (repo *DbSources) DeletePatientAllergy(c *gin.Context, patientId, id string) error {
	_, err := repo.dbConn.Exec(`DELETE FROM patient_allergy WHERE patient_id = $1 AND id = $2`, patientId, id)
	return err
}

func (repo *DbSources) UpdatePatientAllergy(c *gin.Context, p model.PatientAllergy) error {
	updDquery, err := utils.PrepareSQLUpdateStatement(p.Allergy, *p.Allergy.Id)
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
