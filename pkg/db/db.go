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

	CreateUser(c *gin.Context, u model.User) error
	GetUserByEmail(c *gin.Context, u model.User) (model.User, error)
	VerifyUserExists(c *gin.Context, u model.User) error
	GetUserById(c *gin.Context, id string) (model.User, error)
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
