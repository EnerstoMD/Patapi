package controllers

import (
	config "lupus/patapi/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {
	psql, err := config.NewConnect()
	patients, err := psql.GetAllPatients(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
	return
}

func CreatePatient(c *gin.Context) {}
