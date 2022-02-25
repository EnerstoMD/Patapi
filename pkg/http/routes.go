package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, ph PatientHandler) {
	router.GET("/", welcome)
	router.NoRoute(notFound)
	router.GET("/patients", ph.GetAllPatients)
	router.POST("/patients", ph.CreatePatient)
	router.GET("/patients/search", ph.SearchPatientByName)
	router.GET("/patients/:id", ph.GetPatientById)
	router.PATCH("patients/:id", ph.UpdatePatient)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
