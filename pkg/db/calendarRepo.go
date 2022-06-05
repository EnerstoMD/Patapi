package db

import (
	"lupus/patapi/pkg/model"
	"lupus/patapi/utils"

	"github.com/gin-gonic/gin"
)

func (repo *DbSources) DeleteEvent(ctx *gin.Context, id string) error {
	q := `DELETE FROM event WHERE id=` + id
	return repo.execQuery(q)
}

func (repo *DbSources) GetAllEvents(ctx *gin.Context) (events []model.Event, err error) {
	query := "select * from event where created_by=" + ctx.MustGet("userId").(string)
	rows, err := repo.dbConn.Queryx(query)
	for rows.Next() {
		ev := model.Event{}
		err = rows.StructScan(&ev)
		if err != nil {
			return events, err
		}
		// query2 := "select * from patient where id in (select patient_id from event_patientgroup where event_id=" + *ev.Id + ")"
		// rows2, err := repo.dbConn.Queryx(query2)
		// if err != nil {
		// 	return events, err
		// }
		// for rows2.Next() {
		// 	p := model.Patient{}
		// 	err = rows2.StructScan(&p)
		// 	if err != nil {
		// 		return events, err
		// 	}
		// 	ev.ConsultedPatients = append(ev.ConsultedPatients, *p.Id)
		// }
		events = append(events, ev)
	}
	return events, err
}

func (repo *DbSources) CreateEvent(ctx *gin.Context, ev model.Event) error {
	query, err := utils.PrepareSQLInsertStatement(ctx, ev)
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}

func (repo *DbSources) UpdateEvent(c *gin.Context, ev model.Event) error {
	query, err := utils.PrepareSQLUpdateStatement(ev, *ev.Id)
	if err != nil {
		return err
	}
	return repo.execQuery(query)
}

func (repo *DbSources) ConfirmEvent(ctx *gin.Context, id string) error {
	query := `UPDATE event SET is_confirmed=true WHERE id=` + id
	return repo.execQuery(query)
}

func (repo *DbSources) UnconfirmEvent(ctx *gin.Context, id string) error {
	query := `UPDATE event SET is_confirmed=false WHERE id=` + id
	return repo.execQuery(query)
}
