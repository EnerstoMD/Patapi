package db

import (
	"lupus/patapi/pkg/model"

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
