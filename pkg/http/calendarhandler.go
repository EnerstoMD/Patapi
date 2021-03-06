package handler

import (
	"fmt"
	"lupus/patapi/pkg/model"
	calendar "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalendarHandler interface {
	GetAllEvents(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	UpdateEvent(ctx *gin.Context)
	DeleteEvent(ctx *gin.Context)
	ConfirmEvent(ctx *gin.Context)
	UnconfirmEvent(ctx *gin.Context)
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
	createdby := fmt.Sprintf("%v", c.Keys["userId"])
	if createdby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read userId", "error": "userId not found"})
		return
	}
	newEvent.CreatedBy = &createdby
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

func (calendarHandler *calendarHandler) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"status": 400, "msg": "id not filled"})
		return
	}

	err := calendarHandler.calendarService.DeleteEvent(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't delete event-id:" + id, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "Event deleted"})
}

func (calendarHandler *calendarHandler) ConfirmEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"status": 400, "msg": "id not filled"})
		return
	}

	err := calendarHandler.calendarService.ConfirmEvent(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't confirm event-id:" + id, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "Event confirmed"})
}

func (calendarHandler *calendarHandler) UnconfirmEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"status": 400, "msg": "id not filled"})
		return
	}

	err := calendarHandler.calendarService.UnconfirmEvent(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't unconfirm event-id:" + id, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": 204, "msg": "Event unconfirmed"})
}
