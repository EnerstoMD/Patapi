package db

import (
	"fmt"
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"strconv"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) GetAllPatients(ctx *gin.Context) (patients []model.Patient, err error) {
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

func (repo *DbSources) CreatePatient(ctx *gin.Context, patient model.Patient) error {
	query, err := utils.PrepareSQLInsertStatement(ctx, patient)

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

func (repo *DbSources) SearchPatientByName(c *gin.Context, nameOrId string) (patients []model.Patient, err error) {
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

func (repo *DbSources) GetPatientById(c *gin.Context, id string) (patient model.Patient, err error) {
	err = repo.dbConn.Get(&patient, "SELECT * FROM patient WHERE id=$1", id)
	return patient, err
}

func (repo *DbSources) UpdatePatient(c *gin.Context, patient model.Patient) error {
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

func (repo *DbSources) SearchPatientByINSMatricule(c *gin.Context, id string) (patients []model.Patient, err error) {
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

func (repo *DbSources) BatchLoad(c *gin.Context, p []interface{}) error {
	if len(p) == 0 {
		return errors.New("no patient to load")
	}
	queryArgs, err := utils.ReadStructToBeInserted(c, p[0])
	if err != nil {
		return err
	}
	fmt.Println("query: ", `INSERT INTO `+queryArgs[0]+` `+queryArgs[1]+` VALUES `+queryArgs[2])
	_, err = repo.dbConn.NamedExec(`INSERT INTO `+queryArgs[0]+` `+queryArgs[1]+` VALUES `+queryArgs[2], p)
	return err
}

func (repo *DbSources) BatchLoadPatients(c *gin.Context, p []model.Patient) error {
	var err error
	for _, patient := range p {
		if errC := repo.CreatePatient(c, patient); errC != nil {
			err = errors.Wrap(errC, "batch load patients")
		}
	}
	return err
}
