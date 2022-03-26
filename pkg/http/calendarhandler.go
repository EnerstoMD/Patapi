package handler

import (
	calendar "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalendarHandler interface {
	GetAllEvents(ctx *gin.Context)
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
