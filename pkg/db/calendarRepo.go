package db

import (
	"log"
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"

	"github.com/gin-gonic/gin"
)

func (repo *dbRepository) DeleteEvent(ctx *gin.Context, id string) error {
	q := `DELETE FROM event WHERE id=` + id
	log.Println(q)
	return repo.execQuery(q)
}

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
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}

func (repo *dbRepository) UpdateEvent(c *gin.Context, ev model.Event) error {
	query, err := utils.PrepareSQLUpdateStatement(ev, *ev.Id)
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}

func (repo *dbRepository) ConfirmEvent(ctx *gin.Context, id string) error {
	query := `UPDATE event SET is_confirmed=1 WHERE id=` + id
	return repo.execQuery(query)
}
