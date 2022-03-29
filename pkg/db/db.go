package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"os"

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
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, nameOrId string) (patients []model.Patient, err error)

	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
	DeleteEvent(ctx *gin.Context, id string) error
	ConfirmEvent(ctx *gin.Context, id string) error
	UnconfirmEvent(ctx *gin.Context, id string) error
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

func (repo *dbRepository) execQuery(qry string) error {
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

	if _, err = tx.Exec(qry); err != nil {
		return err
	}
	return err
}
