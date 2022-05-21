package handler

import (
	"encoding/csv"
	"fmt"
	"lupus/patapi/pkg/model"
	service "lupus/patapi/pkg/services"
	"lupus/patapi/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	GetAllPatients(ctx *gin.Context)
	CreatePatient(ctx *gin.Context)
	SearchPatientByName(ctx *gin.Context)
	CountSearchPatientByName(c *gin.Context)
	GetPatientById(ctx *gin.Context)
	UpdatePatient(ctx *gin.Context)
	SearchPatientByINSMatricule(c *gin.Context)
	ReadCarteVitale(c *gin.Context)
	BatchLoadPatients(c *gin.Context)
	CreatePatientComment(c *gin.Context)
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
	nameOrId := c.Query("name")
	if nameOrId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "no patient searched"})
		return
	}
	if nameOrId == "*" {
		nameOrId = ""
	}
	pagination := utils.GeneratePaginationFromRequest(c)
	patients, err := patientHandler.patienfileService.SearchPatientByName(c, nameOrId, pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (patientHandler *patientHandler) CountSearchPatientByName(c *gin.Context) {
	nameOrId := c.Query("name")
	if nameOrId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "no patient searched"})
		return
	}
	if nameOrId == "*" {
		nameOrId = ""
	}
	count, err := patientHandler.patienfileService.CountSearchPatientByName(c, nameOrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't count all patients", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, count)
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

func (patientHandler *patientHandler) BatchLoadPatients(c *gin.Context) {
	patientColumnNumber, err := strconv.Atoi(c.Request.FormValue("patientColumnNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patientColumnNumber", "error": err.Error()})
		return
	}
	insColumnNumber, err := strconv.Atoi(c.Request.FormValue("insColumnNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read insColumnNumber", "error": err.Error()})
		return
	}
	if patientColumnNumber < 0 || insColumnNumber < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "csv unreadable"})
		return
	}
	csvFile, _, openErr := c.Request.FormFile("file")
	if openErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't open file", "error": openErr.Error()})
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't read file", "error": err.Error()})
		return
	}
	if len(records) <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "no records found"})
		return
	}

	var patients []model.Patient

	r, _ := regexp.Compile(`^[A-Z\s]+$`)
	for _, record := range records[1:] {
		var patient model.Patient
		var lastname, firstnames string
		var names []string
		var insMatricule string
		for _, namePart := range strings.Split(record[patientColumnNumber-1], " ") {
			if r.MatchString(namePart) {
				lastname += namePart + " "
			} else {
				names = append(names, namePart)
			}
		}
		if len(names) < 1 || lastname == "" {
			continue
		}
		reg, _ := regexp.Compile(`[^0-9]`)
		insMatricule = reg.ReplaceAllString(record[insColumnNumber-1], "")
		lastname = strings.TrimSpace(lastname)
		patient.Lastname = &lastname
		patient.Name = &names[0]
		patient.InsMatricule = &insMatricule
		firstnames = strings.Join(names, ",")
		patient.Firstnames = &firstnames
		patients = append(patients, patient)
	}
	err = patientHandler.patienfileService.BatchLoadPatients(c, patients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't load patients", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": patients})
}

func (patientHandler *patientHandler) CreatePatientComment(c *gin.Context) {
	var comment model.PatientComment
	patId := c.Param("id")

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read comment", "error": err.Error()})
		return
	}
	if comment.PatientId == nil {
		comment.PatientId = &patId
	}
	if comment.PatientId != &patId {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not matching"})
		return
	}
	addedby := fmt.Sprintf("%v", c.Keys["userId"])
	if addedby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read userId"})
		return
	}
	comment.AddedBy = &addedby

	err := patientHandler.patienfileService.CreatePatientComment(c, comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't create patient comment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "patient comment created"})
}
