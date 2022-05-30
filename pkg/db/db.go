package db

import (
	"context"
	"log"
	"lupus/patapi/pkg/model"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DbSources struct {
	dbConn      *sqlx.DB
	redisClient *redis.Client
}
type DbRepository interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, patient model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string, pagination model.Pagination) ([]model.Patient, error)
	CountSearchPatientByName(c *gin.Context, nameOrId string) (int, error)
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, nameOrId string) (patients []model.Patient, err error)
	BatchLoadPatients(c *gin.Context, p []model.Patient) error
	BatchLoad(c *gin.Context, p []interface{}) error

	CreatePatientComment(c *gin.Context, p model.PatientComment) error
	GetPatientComments(c *gin.Context, id string) ([]model.PatientComment, error)
	DeletePatientComment(c *gin.Context, id, commentId string) error

	RegisterPatientDisease(c *gin.Context, p model.PatientDisease) error
	GetPatientDiseases(c *gin.Context, patientId string) ([]model.PatientDisease, error)
	DeletePatientDisease(c *gin.Context, patientId, diseaseId string) error
	UpdatePatientDisease(c *gin.Context, p model.PatientDisease) error

	RegisterPatientAllergy(c *gin.Context, p model.PatientAllergy) error
	GetPatientAllergies(c *gin.Context, patientId string) ([]model.PatientAllergy, error)
	DeletePatientAllergy(c *gin.Context, patientId, allergyId string) error
	UpdatePatientAllergy(c *gin.Context, p model.PatientAllergy) error

	RegisterPatienTreatment(c *gin.Context, p model.PatientTreatment) error
	GetPatientTreatments(c *gin.Context, patientId string) ([]model.PatientTreatment, error)
	DeletePatientTreatment(c *gin.Context, patientId, treatId string) error
	UpdatePatientTreatment(c *gin.Context, p model.PatientTreatment) error

	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
	DeleteEvent(ctx *gin.Context, id string) error
	ConfirmEvent(ctx *gin.Context, id string) error
	UnconfirmEvent(ctx *gin.Context, id string) error

	CreateUser(c *gin.Context, u model.User) error
	GetUserByEmail(c *gin.Context, u model.User) (model.User, error)
	VerifyUserExists(c *gin.Context, u model.User) error
	GetUserById(c *gin.Context, id string) (model.User, error)
	GetUsers(c *gin.Context) ([]model.User, error)
	UpdateUser(c *gin.Context, u model.User) error

	GetUserRoles(c *gin.Context, id string) ([]int, error)
}

type TokenRepository interface {
	SetRefreshToken(c *gin.Context, userID, tokenID string, expiresIn time.Duration) error
	ValidateToken(c *gin.Context, userID, previoustokenID string) error
}

func NewDbConnect() *DbSources {
	dbURL := "postgres://" + os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") + "@" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + "/" + os.Getenv("DBNAME")
	conn, err := sqlx.Connect("pgx", dbURL)
	if conn == nil || err != nil {
		log.Fatalf("Failed to connect to db")
		os.Exit(100)
	}
	log.Printf("Connected to Postgres")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis")
		os.Exit(100)
	}
	log.Printf("Connected to Redis")
	//defer conn.Close()
	return &DbSources{
		dbConn:      conn,
		redisClient: redisClient,
	}
}

func (repo *DbSources) execQuery(qry string) error {
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
