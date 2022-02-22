package controllers

import (
	"log"
	config "lupus/patapi/configs"
	"lupus/patapi/models"
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

func CreatePatient(c *gin.Context) {
	var newPatient models.Patient
	if err := c.BindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient", "error": err.Error()})
		return
	}
	log.Println(newPatient)
	psql, err := config.NewConnect()
	err = psql.CreatePatient(c, newPatient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't insert patient into db", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "insert OK"})
}
