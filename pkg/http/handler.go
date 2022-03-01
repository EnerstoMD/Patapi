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
	SearchPatientByName(ctx *gin.Context)
	GetPatientById(ctx *gin.Context)
	UpdatePatient(ctx *gin.Context)
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
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "Patient registered"})
}

func (patientHandler *patientHandler) SearchPatientByName(c *gin.Context) {
	// var searchedPatient model.Patient
	// if err := c.BindJSON(&searchedPatient); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient", "error": err.Error()})
	// 	return
	// }
	nameOrId := c.Query("name")
	//var nameOrId string

	// switch {
	// case *searchedPatient.Name != "":
	// 	nameOrId = *searchedPatient.Name
	// case *searchedPatient.Lastname != "":
	// 	nameOrId = *searchedPatient.Lastname
	// case *searchedPatient.InsMatricule != "":
	// 	nameOrId = *searchedPatient.InsMatricule
	// }
	patients, err := patientHandler.patienfileService.SearchPatientByName(c, nameOrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
	return
}

func (patientHandler *patientHandler) GetPatientById(ctx *gin.Context) {
	id := ctx.Param("id")
	patient, err := patientHandler.patienfileService.GetPatientById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get this patient", "error": err.Error()})
		return
	}
	if patient.Id == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": 404, "msg": "Patient not found"})
		return
	}
	ctx.JSON(http.StatusOK, patient)
	return
}

func (ph *patientHandler) UpdatePatient(ctx *gin.Context) {
	id := ctx.Param("id")
	var patientToUpdate model.Patient
	if err := ctx.BindJSON(&patientToUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient", "error": err.Error()})
		return
	}
	if patientToUpdate.Id == nil {
		patientToUpdate.Id = &id
	}
	err := ph.patienfileService.UpdatePatient(ctx, patientToUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't insert patient into db", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "msg": "Patient updated"})
}
