package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"

	"github.com/gin-gonic/gin"
)

func (repo *dbRepository) GetAllEvents(ctx *gin.Context) (events []model.Event, err error) {
	query := "select * from event"
	rows, err := repo.dbConn.Queryx(query)
	for rows.Next() {
		ev := model.Event{}
		err = rows.StructScan(&ev)
		events = append(events, ev)
	}
	return events, err
}

func (repo *dbRepository) CreateEvent(ctx *gin.Context, ev model.Event) error {
	query, err := utils.PrepareSQLInsertStatement(ev)
	log.Println(query)
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

func (repo *dbRepository) UpdateEvent(c *gin.Context, ev model.Event) error {
	query, err := utils.PrepareSQLUpdateStatement(ev, *ev.Id)
	log.Println(query)
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
