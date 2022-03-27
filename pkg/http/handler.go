package handler

import (
	"log"
	"lupus/patapi/pkg/model"
	service "lupus/patapi/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	GetAllPatients(ctx *gin.Context)
	CreatePatient(ctx *gin.Context)
	SearchPatientByName(ctx *gin.Context)
	GetPatientById(ctx *gin.Context)
	UpdatePatient(ctx *gin.Context)
	SearchPatientByINSMatricule(c *gin.Context)
	ReadCarteVitale(c *gin.Context)
}

type patientHandler struct {
	patienfileService service.PatService
}

func NewPatientHandler(patientService service.PatService) PatientHandler {
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
	log.Println("nameOrId: ", nameOrId)
	if nameOrId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "no patient searched"})
		return
	}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't update patient into db", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "msg": "Patient updated"})
}

func (patientHandler *patientHandler) SearchPatientByINSMatricule(c *gin.Context) {
	id := c.Query("ins")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "no patient searched"})
		return
	}

	patients, err := patientHandler.patienfileService.SearchPatientByINSMatricule(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	if len(patients) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient found"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (patientHandler *patientHandler) ReadCarteVitale(c *gin.Context) {
	var searchedP model.CardPeek

	if err := c.BindXML(&searchedP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read xml", "error": err.Error()})
		return
	}
	patient, err := patientHandler.patienfileService.ReadCarteVitale(c, searchedP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}
