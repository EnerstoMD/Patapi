package config

import (
	"context"
	"fmt"
	"log"
	"lupus/patapi/models"
	"lupus/patapi/utils"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgreSQL struct {
	dbConn *sqlx.DB
}

func NewConnect() (*PostgreSQL, error) {
	dbURL := "postgres://" + os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") + "@" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + "/" + os.Getenv("DBNAME")
	conn, err := sqlx.Open("pgx", dbURL)
	if conn == nil || err != nil {
		log.Fatalf("Failed to connect to db")
		os.Exit(100)
	}
	log.Printf("Connected to DB")
	return &PostgreSQL{dbConn: conn}, nil
}

func (p *PostgreSQL) GetAllPatients(ctx context.Context) (patients []models.Patient, err error) {
	query := fmt.Sprintf("select * from patient")
	defer p.dbConn.Close()

	rows, err := p.dbConn.Queryx(query)
	defer rows.Close()
	for rows.Next() {
		var patient models.Patient
		err = rows.StructScan(&patient)
		patients = append(patients, patient)
	}
	return patients, err
}

func (p *PostgreSQL) CreatePatient(ctx context.Context, patient models.Patient) error {
	query, err := utils.PrepareSQLInsertStatement(patient)
	if err != nil {
		return err
	}
	tx, err := p.dbConn.Begin()
	_, err = tx.Exec(query)
	err = tx.Commit()
	return err
}
