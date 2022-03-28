package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, ph PatientHandler, ch CalendarHandler) {
	router.GET("/", welcome)
	router.NoRoute(notFound)
	router.Use(cors.Default())

	router.GET("v1/patients", ph.GetAllPatients)
	router.POST("v1/patients", ph.CreatePatient)
	router.GET("v1/patients/search", ph.SearchPatientByName)
	router.GET("v1/patients/:id", ph.GetPatientById)
	router.PATCH("v1/patients/:id", ph.UpdatePatient)
	router.GET("v1/patients/ins", ph.SearchPatientByINSMatricule)
	router.GET("v1/patients/card", ph.ReadCarteVitale)

	router.GET("v1/calendar", ch.GetAllEvents)
	router.POST("v1/calendar", ch.CreateEvent)
	router.PATCH("v1/calendar/:id", ch.UpdateEvent)
	router.DELETE("v1/calendar/:id", ch.DeleteEvent)
	router.PATCH("v1/calendar/:id/confirm", ch.ConfirmEvent)

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
