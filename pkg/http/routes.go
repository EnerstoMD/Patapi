package handler

import (
	"lupus/patapi/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, ph PatientHandler, ch CalendarHandler, uh UserHandler) {
	router.GET("", welcome)
	router.NoRoute(notFound)
	v1 := router.Group("v1")
	{
		user := v1.Group("user")
		{
			user.POST("register", uh.Register)
			user.POST("login", uh.Login)
		}

		patient := v1.Group("patient")
		patient.Use(middleware.BearerAuth())
		{
			patient.GET("", ph.GetAllPatients)
			patient.POST("", ph.CreatePatient)
			patient.GET("search", ph.SearchPatientByName)
			patient.GET(":id", ph.GetPatientById)
			patient.PATCH(":id", ph.UpdatePatient)
			patient.GET("ins", ph.SearchPatientByINSMatricule)
			patient.GET("card", ph.ReadCarteVitale)
		}

		calendar := v1.Group("calendar")
		calendar.Use(middleware.BearerAuth())
		{
			calendar.GET("", ch.GetAllEvents)
			calendar.POST("", ch.CreateEvent)
			calendar.PATCH(":id", ch.UpdateEvent)
			calendar.DELETE(":id", ch.DeleteEvent)
			calendar.PATCH(":id/confirm", ch.ConfirmEvent)
			calendar.PATCH(":id/unconfirm", ch.UnconfirmEvent)
		}

	}

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To PATIENTAPI",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
