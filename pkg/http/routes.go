package handler

import (
	"lupus/patapi/pkg/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, ph PatientHandler, ch CalendarHandler, uh UserHandler) {
	router.GET("/", welcome)
	router.NoRoute(notFound)
	router.Use(cors.Default())

	router.POST("v1/user/register", uh.Register)
	router.POST("v1/user/login", uh.Login)

	pat := router.Group("v1/patients")
	pat.Use(middleware.BearerAuth())
	{
		pat.GET("/", ph.GetAllPatients)
		pat.POST("/", ph.CreatePatient)
		pat.GET("/search", ph.SearchPatientByName)
		pat.GET("/:id", ph.GetPatientById)
		pat.PATCH("/:id", ph.UpdatePatient)
		pat.GET("/ins", ph.SearchPatientByINSMatricule)
		pat.GET("/card", ph.ReadCarteVitale)
	}

	router.GET("v1/calendar", ch.GetAllEvents)
	router.POST("v1/calendar", ch.CreateEvent)
	router.PATCH("v1/calendar/:id", ch.UpdateEvent)
	router.DELETE("v1/calendar/:id", ch.DeleteEvent)
	router.PATCH("v1/calendar/:id/confirm", ch.ConfirmEvent)
	router.PATCH("v1/calendar/:id/unconfirm", ch.UnconfirmEvent)

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
