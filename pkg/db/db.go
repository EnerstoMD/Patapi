package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type dbRepository struct {
	dbConn *sqlx.DB
}
type DbRepository interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, patient model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string) ([]model.Patient, error)
}

func NewDbConnect() *dbRepository {
	dbURL := "postgres://" + os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") + "@" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + "/" + os.Getenv("DBNAME")
	conn, err := sqlx.Open("pgx", dbURL)
	if conn == nil || err != nil {
		log.Fatalf("Failed to connect to db")
		os.Exit(100)
	}
	log.Printf("Connected to DB")
	//defer conn.Close()
	return &dbRepository{dbConn: conn}
}

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
	for rows.Next() {
		patient := model.Patient{}
		err = rows.StructScan(&patient)
		patients = append(patients, patient)
	}
	return patients, err
}
