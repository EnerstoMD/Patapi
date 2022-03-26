package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (repo *dbRepository) GetAllPatients(ctx *gin.Context) (patients []model.Patient, err error) {
	query := "select * from patient"
	//defer repo.dbConn.Close()
	rows, err := repo.dbConn.Queryx(query)
	//defer rows.Close()
	for rows.Next() {
		patient := model.Patient{}
		err = rows.StructScan(&patient)
		patients = append(patients, patient)
	}
	return patients, err
}

func (repo *dbRepository) CreatePatient(ctx *gin.Context, patient model.Patient) error {
	query, err := utils.PrepareSQLInsertStatement(patient)

	if err != nil {
		return err
	}
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

func (repo *dbRepository) SearchPatientByName(c *gin.Context, nameOrId string) (patients []model.Patient, err error) {
	query := `SELECT * FROM patient WHERE (name LIKE '%` + nameOrId + `%' OR lastname LIKE '%` + nameOrId + `%')`
	if numSecu, err := strconv.Atoi(nameOrId); err == nil && numSecu > 99999999999999 {
		query = `SELECT * FROM patient WHERE ins_matricule=` + nameOrId
	}

	rows, err := repo.dbConn.Queryx(query)
	//rows, err := repo.dbConn.NamedQuery(`SELECT * FROM patient WHERE name =:nameOrId`, nameOrId)
	//defer rows.Close()
	log.Println("query: ", query)
	for rows.Next() {
		patient := model.Patient{}
		err = rows.StructScan(&patient)
		patients = append(patients, patient)
	}
	return patients, err
}

func (repo *dbRepository) GetPatientById(c *gin.Context, id string) (patient model.Patient, err error) {
	err = repo.dbConn.Get(&patient, "SELECT * FROM patient WHERE id=$1", id)
	return patient, err
}

func (repo *dbRepository) UpdatePatient(c *gin.Context, patient model.Patient) error {
	query, err := utils.PrepareSQLUpdateStatement(patient, *patient.Id)

	if err != nil {
		return err
	}
	tx, err := repo.dbConn.Begin()
	if err != nil {
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

func (repo *dbRepository) SearchPatientByINSMatricule(c *gin.Context, id string) (patients []model.Patient, err error) {
	// numSecu, err := strconv.Atoi(id)
	// if err != nil {
	// 	return nil, err
	// }

	query := `SELECT * FROM patient WHERE ins_matricule LIKE '` + id + `'`
	rows, err := repo.dbConn.Queryx(query)
	// rows, err := repo.dbConn.Queryx(query)
	for rows.Next() {
		patient := model.Patient{}
		err = rows.StructScan(&patient)
		patients = append(patients, patient)
	}
	return patients, err
}
