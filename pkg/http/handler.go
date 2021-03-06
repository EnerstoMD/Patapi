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
	GetPatientComments(c *gin.Context)
	DeletePatientComment(c *gin.Context)

	RegisterPatientDisease(c *gin.Context)
	GetPatientDiseases(c *gin.Context)
	DeletePatientDisease(c *gin.Context)
	UpdatePatientDisease(c *gin.Context)

	RegisterPatientAllergy(c *gin.Context)
	GetPatientAllergies(c *gin.Context)
	DeletePatientAllergy(c *gin.Context)
	UpdatePatientAllergy(c *gin.Context)

	RegisterPatientTreatment(c *gin.Context)
	GetPatientTreatments(c *gin.Context)
	DeletePatientTreatment(c *gin.Context)
	UpdatePatientTreatment(c *gin.Context)

	RegisterPatientHistory(c *gin.Context)
	GetPatientHistory(c *gin.Context)
	DeletePatientHistory(c *gin.Context)
	UpdatePatientHistory(c *gin.Context)

	RegisterHospitalisation(c *gin.Context)
	GetHospitalisations(c *gin.Context)
	DeleteHospitalisation(c *gin.Context)
	UpdateHospitalisation(c *gin.Context)
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

func (patientHandler *patientHandler) GetPatientComments(c *gin.Context) {
	patId := c.Param("id")
	comments, err := patientHandler.patienfileService.GetPatientComments(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get patient comments", "error": err.Error()})
		return
	}
	if len(comments) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient comments found"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (patientHandler *patientHandler) DeletePatientComment(c *gin.Context) {
	patId := c.Param("id")
	commentId := c.Param("commentid")
	err := patientHandler.patienfileService.DeletePatientComment(c, patId, commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete patient comment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient comment deleted"})
}

func (patientHandler *patientHandler) RegisterPatientDisease(c *gin.Context) {
	var patientDisease model.PatientDisease
	if err := c.BindJSON(&patientDisease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient disease", "error": err.Error()})
		return
	}

	if patientDisease.PatientId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not set"})
		return
	}
	addedby := fmt.Sprintf("%v", c.Keys["userId"])
	if addedby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read userId"})
		return
	}
	patientDisease.AddedBy = &addedby
	err := patientHandler.patienfileService.RegisterPatientDisease(c, patientDisease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't register patient disease", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "patient disease registered"})
}

func (patientHandler *patientHandler) GetPatientDiseases(c *gin.Context) {
	patId := c.Param("id")
	diseases, err := patientHandler.patienfileService.GetPatientDiseases(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get patient diseases", "error": err.Error()})
		return
	}
	if len(diseases) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient diseases found"})
		return
	}
	c.JSON(http.StatusOK, diseases)
}

func (patientHandler *patientHandler) DeletePatientDisease(c *gin.Context) {
	patId := c.Param("id")
	diseaseId := c.Param("patdiseaseid")
	err := patientHandler.patienfileService.DeletePatientDisease(c, patId, diseaseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete patient disease", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient disease deleted"})
}

func (patientHandler *patientHandler) UpdatePatientDisease(c *gin.Context) {
	var patientDisease model.PatientDisease
	patId := c.Param("id")
	diseaseId := c.Param("diseaseid")
	if err := c.BindJSON(&patientDisease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient disease", "error": err.Error()})
		return
	}
	if patientDisease.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient disease id not set"})
		return
	}
	if patientDisease.PatientId == nil {
		patientDisease.PatientId = &patId
	}
	if patientDisease.Disease.Id == nil {
		patientDisease.Disease.Id = &diseaseId
	}
	if patId != *patientDisease.PatientId || diseaseId != *patientDisease.Disease.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id or disease id not matching"})
		return
	}
	err := patientHandler.patienfileService.UpdatePatientDisease(c, patientDisease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't update patient disease", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient disease updated"})
}

func (patientHandler *patientHandler) RegisterPatientAllergy(c *gin.Context) {
	var patientAllergy model.PatientAllergy
	if err := c.BindJSON(&patientAllergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient allergy", "error": err.Error()})
		return
	}

	if patientAllergy.PatientId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not set"})
		return
	}
	addedby := fmt.Sprintf("%v", c.Keys["userId"])
	if addedby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read userId"})
		return
	}
	patientAllergy.AddedBy = &addedby
	err := patientHandler.patienfileService.RegisterPatientAllergy(c, patientAllergy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't register patient allergy", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "patient disease registered"})
}

func (patientHandler *patientHandler) GetPatientAllergies(c *gin.Context) {
	patId := c.Param("id")
	allergies, err := patientHandler.patienfileService.GetPatientAllergies(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get patient allergies", "error": err.Error()})
		return
	}
	if len(allergies) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient allergies found"})
		return
	}
	c.JSON(http.StatusOK, allergies)
}

func (patientHandler *patientHandler) DeletePatientAllergy(c *gin.Context) {
	patId := c.Param("id")
	patallergyId := c.Param("patallergyid")
	err := patientHandler.patienfileService.DeletePatientAllergy(c, patId, patallergyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete patient allergy", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient allergy deleted"})
}

func (patientHandler *patientHandler) UpdatePatientAllergy(c *gin.Context) {
	var patientAllergy model.PatientAllergy
	patId := c.Param("id")
	diseaseId := c.Param("allergyid")
	if err := c.BindJSON(&patientAllergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient allergy", "error": err.Error()})
		return
	}
	if patientAllergy.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient allergy id not set"})
		return
	}
	if patientAllergy.PatientId == nil {
		patientAllergy.PatientId = &patId
	}
	if patientAllergy.Allergy.Id == nil {
		patientAllergy.Allergy.Id = &diseaseId
	}
	if patId != *patientAllergy.PatientId || diseaseId != *patientAllergy.Allergy.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id or allergy id not matching"})
		return
	}
	err := patientHandler.patienfileService.UpdatePatientAllergy(c, patientAllergy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't update patient allergy", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient allergy updated"})
}

func (patientHandler *patientHandler) RegisterPatientTreatment(c *gin.Context) {
	var patientTreatment model.PatientTreatment
	if err := c.BindJSON(&patientTreatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient treatment", "error": err.Error()})
		return
	}

	if patientTreatment.PatientId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not set"})
		return
	}
	addedby := fmt.Sprintf("%v", c.Keys["userId"])
	if addedby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read userId"})
		return
	}
	patientTreatment.AddedBy = &addedby
	err := patientHandler.patienfileService.RegisterPatientTreatment(c, patientTreatment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't register patient treatment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "patient treatment registered"})
}

func (patientHandler *patientHandler) GetPatientTreatments(c *gin.Context) {
	patId := c.Param("id")
	treatments, err := patientHandler.patienfileService.GetPatientTreatments(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get patient treaments", "error": err.Error()})
		return
	}
	if len(treatments) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient treatments found"})
		return
	}
	c.JSON(http.StatusOK, treatments)
}

func (patientHandler *patientHandler) DeletePatientTreatment(c *gin.Context) {
	patId := c.Param("id")
	pattreatmentId := c.Param("pattreatmentid")
	err := patientHandler.patienfileService.DeletePatientTreatment(c, patId, pattreatmentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete patient treatment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient treatment deleted"})
}

func (patientHandler *patientHandler) UpdatePatientTreatment(c *gin.Context) {
	var patientTreatment model.PatientTreatment
	patId := c.Param("id")
	diseaseId := c.Param("treatmentid")
	if err := c.BindJSON(&patientTreatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient treatment", "error": err.Error()})
		return
	}
	if patientTreatment.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient allergy id not set"})
		return
	}
	if patientTreatment.PatientId == nil {
		patientTreatment.PatientId = &patId
	}
	if patientTreatment.Treatment.Id == nil {
		patientTreatment.Treatment.Id = &diseaseId
	}
	if patId != *patientTreatment.PatientId || diseaseId != *patientTreatment.Treatment.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id or treatment id not matching"})
		return
	}
	err := patientHandler.patienfileService.UpdatePatientTreatment(c, patientTreatment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't update patient treatment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient treatment updated"})
}

func (patientHandler *patientHandler) RegisterPatientHistory(c *gin.Context) {
	var patientHistory model.PatientHistory
	if err := c.BindJSON(&patientHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient history", "error": err.Error()})
		return
	}

	if patientHistory.PatientId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not set"})
		return
	}

	err := patientHandler.patienfileService.RegisterPatientHistory(c, patientHistory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't register patient history", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "patient history registered"})
}

func (patientHandler *patientHandler) GetPatientHistory(c *gin.Context) {
	patId := c.Param("id")
	histories, err := patientHandler.patienfileService.GetPatientHistory(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get patient history", "error": err.Error()})
		return
	}
	if len(histories) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no patient history found"})
		return
	}
	c.JSON(http.StatusOK, histories)
}

func (patientHandler *patientHandler) DeletePatientHistory(c *gin.Context) {
	patId := c.Param("id")
	patHistoryId := c.Param("patHistoryId")
	err := patientHandler.patienfileService.DeletePatientHistory(c, patId, patHistoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete patient history", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient history deleted"})
}

func (patientHandler *patientHandler) UpdatePatientHistory(c *gin.Context) {
	var patientHistory model.PatientHistory
	patId := c.Param("id")
	historyId := c.Param("historyid")
	if err := c.BindJSON(&patientHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read patient history", "error": err.Error()})
		return
	}
	if patientHistory.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient history id not set"})
		return
	}
	if patientHistory.PatientId == nil {
		patientHistory.PatientId = &patId
	}
	if patId != *patientHistory.PatientId || historyId != *patientHistory.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not matching"})
		return
	}
	err := patientHandler.patienfileService.UpdatePatientHistory(c, patientHistory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't update patient history", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "patient history updated"})
}

func (patientHandler *patientHandler) RegisterHospitalisation(c *gin.Context) {
	var hospitalisation model.Hospitalisation
	patId := c.Param("id")
	if err := c.BindJSON(&hospitalisation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read hospitalisation", "error": err.Error()})
		return
	}

	if hospitalisation.PatientId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not set"})
		return
	}

	if patId != *hospitalisation.PatientId {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient id not matching"})
		return
	}

	err := patientHandler.patienfileService.RegisterHospitalisation(c, hospitalisation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't register hospitalisation", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 201, "msg": "hospitalisation registered"})
}

func (patientHandler *patientHandler) GetHospitalisations(c *gin.Context) {
	patId := c.Param("id")
	hospitalisations, err := patientHandler.patienfileService.GetHospitalisations(c, patId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't get hospitalisation", "error": err.Error()})
		return
	}
	if len(hospitalisations) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": 204, "message": "no hospitalisation found"})
		return
	}
	c.JSON(http.StatusOK, hospitalisations)
}

func (patientHandler *patientHandler) DeleteHospitalisation(c *gin.Context) {
	patId := c.Param("id")
	hospitalisationId := c.Param("patHospitalisationId")
	err := patientHandler.patienfileService.DeleteHospitalisation(c, patId, hospitalisationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't delete hospitalisation", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "hospitalisation deleted"})
}

func (patientHandler *patientHandler) UpdateHospitalisation(c *gin.Context) {
	var hospitalisation model.Hospitalisation
	patId := c.Param("id")
	hospitalisationId := c.Param("hospitalisationid")
	if err := c.BindJSON(&hospitalisation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "can't read hospitalisation", "error": err.Error()})
		return
	}
	if hospitalisation.Id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "hospitalisation id not set"})
		return
	}
	if hospitalisation.PatientId == nil {
		hospitalisation.PatientId = &patId
	}
	if patId != *hospitalisation.PatientId || hospitalisationId != *hospitalisation.Id {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "msg": "patient or hospitalisation ids not matching"})
		return
	}
	err := patientHandler.patienfileService.UpdateHospitalisation(c, hospitalisation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "can't update hospitalisation", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "hospitalisation updated"})
}
