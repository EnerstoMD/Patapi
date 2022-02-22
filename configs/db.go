package config

import (
	"context"
	"fmt"
	"log"
	"lupus/patapi/models"
	"os"

	"github.com/jackc/pgx/v4"
)

type PostgreSQL struct {
	dbConn *pgx.Conn
}

func NewConnect() (*PostgreSQL, error) {
	dbURL := "postgres://" + os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") + "@" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + "/" + os.Getenv("DBNAME")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if conn == nil || err != nil {
		log.Fatalf("Failed to connect to db")
		os.Exit(100)
	}
	log.Printf("Connected to DB")
	return &PostgreSQL{dbConn: conn}, nil
}

func (p *PostgreSQL) GetAllPatients(ctx context.Context) (patients []models.Patient, err error) {
	query := fmt.Sprintf("select name from people")
	log.Println("query:", query)
	defer p.dbConn.Close(ctx)

	rows, err := p.dbConn.Query(ctx, query)
	defer rows.Close()
	for rows.Next() {
		val, _ := rows.Values()
		var patient models.Patient
		patient.Name = val[0].(string)
		patients = append(patients, patient)
	}
	return patients, err
}
