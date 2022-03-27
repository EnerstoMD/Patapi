package handler

import (
	"lupus/patapi/pkg/model"
	calendar "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalendarHandler interface {
	GetAllEvents(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	UpdateEvent(ctx *gin.Context)
}

type calendarHandler struct {
	calendarService calendar.CalService
}

func NewCalendarHandler(calService calendar.CalService) CalendarHandler {
	return &calendarHandler{
		calendarService: calService,
	}
}

func (calendarHandler *calendarHandler) GetAllEvents(c *gin.Context) {
	calendar, err := calendarHandler.calendarService.GetAllEvents(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all events", "error": err.Error()})
	}
	c.JSON(http.StatusOK, calendar)
}

func (calendarHandler *calendarHandler) CreateEvent(c *gin.Context) {
	var newEvent model.Event
	if err := c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read event", "error": err.Error()})
		return
	}
	err := calendarHandler.calendarService.CreateEvent(c, newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't insert event into db", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "Event registered"})
}

func (calendarHandler *calendarHandler) UpdateEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	var eventToUpt model.Event
	if err := ctx.BindJSON(&eventToUpt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient", "error": err.Error()})
		return
	}
	if eventToUpt.Id == nil {
		eventToUpt.Id = &id
	}
	err := calendarHandler.calendarService.UpdateEvent(ctx, eventToUpt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't update event into db", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "msg": "Event updated"})

}
