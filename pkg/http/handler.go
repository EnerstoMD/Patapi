package handler

import (
	"lupus/patapi/pkg/model"
	patientfile "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	GetAllPatients(ctx *gin.Context)
	CreatePatient(ctx *gin.Context)
	GetPatientByName(ctx *gin.Context)
}

type patientHandler struct {
	patienfileService patientfile.Service
}

func NewHandler(patientService patientfile.Service) PatientHandler {
	return &patientHandler{
		patienfileService: patientService,
	}
}

func (patientHandler *patientHandler) GetAllPatients(c *gin.Context) {
	patients, err := patientHandler.patienfileService.GetAllPatients(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
	return
}

func (patientHandler *patientHandler) CreatePatient(c *gin.Context) {
	var newPatient model.Patient
	if err := c.BindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient", "error": err.Error()})
		return
	}

	err := patientHandler.patienfileService.CreatePatient(c, newPatient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't insert patient into db", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "insert OK"})
}

func (patientHandler *patientHandler) GetPatientByName(c *gin.Context) {
	nameOrId := c.Param("name")
	patients, err := patientHandler.patienfileService.GetPatientByName(c, nameOrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
	return
}
