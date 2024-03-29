package service

import (
	"errors"
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type PatService interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string, pagination model.Pagination) ([]model.Patient, error)
	CountSearchPatientByName(c *gin.Context, nameOrId string) (int, error)
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error)
	ReadCarteVitale(c *gin.Context, p model.CardPeek) (model.Patient, error)
	BatchLoadPatients(c *gin.Context, p []model.Patient) error

	CreatePatientComment(c *gin.Context, p model.PatientComment) error
	GetPatientComments(c *gin.Context, id string) ([]model.PatientComment, error)
	DeletePatientComment(c *gin.Context, id, commentId string) error

	RegisterPatientDisease(c *gin.Context, p model.PatientDisease) error
	GetPatientDiseases(c *gin.Context, patId string) ([]model.PatientDisease, error)
	DeletePatientDisease(c *gin.Context, patId, diseaseId string) error
	UpdatePatientDisease(c *gin.Context, patientDisease model.PatientDisease) error

	RegisterPatientAllergy(c *gin.Context, p model.PatientAllergy) error
	GetPatientAllergies(c *gin.Context, patId string) ([]model.PatientAllergy, error)
	DeletePatientAllergy(c *gin.Context, patId, allergyId string) error
	UpdatePatientAllergy(c *gin.Context, patientAllergy model.PatientAllergy) error

	RegisterPatientTreatment(c *gin.Context, p model.PatientTreatment) error
	GetPatientTreatments(c *gin.Context, patId string) ([]model.PatientTreatment, error)
	DeletePatientTreatment(c *gin.Context, patId, treatmentId string) error
	UpdatePatientTreatment(c *gin.Context, patientTreatment model.PatientTreatment) error

	RegisterPatientHistory(c *gin.Context, p model.PatientHistory) error
	GetPatientHistory(c *gin.Context, patId string) ([]model.PatientHistory, error)
	DeletePatientHistory(c *gin.Context, patId, historyId string) error
	UpdatePatientHistory(c *gin.Context, patientHistory model.PatientHistory) error

	RegisterHospitalisation(c *gin.Context, p model.Hospitalisation) error
	GetHospitalisations(c *gin.Context, patId string) ([]model.Hospitalisation, error)
	DeleteHospitalisation(c *gin.Context, patId, hospitalisationId string) error
	UpdateHospitalisation(c *gin.Context, hospitalisation model.Hospitalisation) error
}
type Db interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string, pagination model.Pagination) ([]model.Patient, error)
	CountSearchPatientByName(c *gin.Context, nameOrId string) (int, error)
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error)
	BatchLoadPatients(c *gin.Context, p []model.Patient) error

	CreatePatientComment(c *gin.Context, p model.PatientComment) error
	GetPatientComments(c *gin.Context, id string) ([]model.PatientComment, error)
	DeletePatientComment(c *gin.Context, id, commentId string) error

	RegisterPatientDisease(c *gin.Context, p model.PatientDisease) error
	GetPatientDiseases(c *gin.Context, patId string) ([]model.PatientDisease, error)
	DeletePatientDisease(c *gin.Context, patId, diseaseId string) error
	UpdatePatientDisease(c *gin.Context, patientDisease model.PatientDisease) error

	RegisterPatientAllergy(c *gin.Context, p model.PatientAllergy) error
	GetPatientAllergies(c *gin.Context, patId string) ([]model.PatientAllergy, error)
	DeletePatientAllergy(c *gin.Context, patId, allergyId string) error
	UpdatePatientAllergy(c *gin.Context, patientAllergy model.PatientAllergy) error

	RegisterPatientTreatment(c *gin.Context, p model.PatientTreatment) error
	GetPatientTreatments(c *gin.Context, patId string) ([]model.PatientTreatment, error)
	DeletePatientTreatment(c *gin.Context, patId, treatmentId string) error
	UpdatePatientTreatment(c *gin.Context, patientTreatment model.PatientTreatment) error

	RegisterPatientHistory(c *gin.Context, p model.PatientHistory) error
	GetPatientHistory(c *gin.Context, patId string) ([]model.PatientHistory, error)
	DeletePatientHistory(c *gin.Context, patId, historyId string) error
	UpdatePatientHistory(c *gin.Context, patientHistory model.PatientHistory) error

	RegisterHospitalisation(c *gin.Context, p model.Hospitalisation) error
	GetHospitalisations(c *gin.Context, patId string) ([]model.Hospitalisation, error)
	DeleteHospitalisation(c *gin.Context, patId, hospId string) error
	UpdateHospitalisation(c *gin.Context, hosp model.Hospitalisation) error
}

type patService struct {
	d Db
}

func NewPatService(d Db) PatService {
	return &patService{d}
}

func (s *patService) GetAllPatients(c *gin.Context) ([]model.Patient, error) {
	return s.d.GetAllPatients(c)
}

func (s *patService) CreatePatient(c *gin.Context, p model.Patient) error {
	return s.d.CreatePatient(c, p)
}

func (s *patService) SearchPatientByName(c *gin.Context, nameOrId string, pagination model.Pagination) ([]model.Patient, error) {
	return s.d.SearchPatientByName(c, nameOrId, pagination)
}

func (s *patService) CountSearchPatientByName(c *gin.Context, nameOrId string) (int, error) {
	return s.d.CountSearchPatientByName(c, nameOrId)
}

func (s *patService) GetPatientById(c *gin.Context, id string) (model.Patient, error) {
	return s.d.GetPatientById(c, id)
}

func (s *patService) UpdatePatient(c *gin.Context, p model.Patient) error {
	return s.d.UpdatePatient(c, p)
}

func (s *patService) SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error) {
	return s.d.SearchPatientByINSMatricule(c, id)
}

func (s *patService) BatchLoadPatients(c *gin.Context, p []model.Patient) error {
	return s.d.BatchLoadPatients(c, p)
}

func (s *patService) ReadCarteVitale(c *gin.Context, cp model.CardPeek) (patient model.Patient, e error) {
	var nom, prenom, nir string
	for i := 0; i < len(cp.Node.Node); i++ {
		for _, node1 := range cp.Node.Node {
			for _, node2 := range node1.Node {
				for _, node3 := range node2.Node {
					for _, node4 := range node3.Node {
						for _, node5 := range node4.Node {
							for k, attr5 := range node5.Attr {
								if attr5.Name == "label" {
									switch attr5.Text {
									case "Nom":
										nom = node5.Attr[k+4].Text
									case "Prénom":
										prenom = node5.Attr[k+4].Text
									case "Numéro de sécurité sociale":
										nir = node5.Attr[k+4].Text
									default:
										continue
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if prenom == "" || nom == "" || nir == "" {
		e = errors.New("can't extract prenom:" + prenom + ", nom:" + nom + " ,nir:" + nir)
		return patient, e
	}
	patient.Name = &prenom
	patient.Lastname = &nom
	patient.InsMatricule = &nir
	//if no id, not in db
	return patient, e
}

func (s *patService) CreatePatientComment(c *gin.Context, p model.PatientComment) error {
	return s.d.CreatePatientComment(c, p)
}

func (s *patService) GetPatientComments(c *gin.Context, id string) ([]model.PatientComment, error) {
	return s.d.GetPatientComments(c, id)
}

func (s *patService) DeletePatientComment(c *gin.Context, id, commentId string) error {
	return s.d.DeletePatientComment(c, id, commentId)
}

func (s *patService) RegisterPatientDisease(c *gin.Context, p model.PatientDisease) error {
	return s.d.RegisterPatientDisease(c, p)
}

func (s *patService) GetPatientDiseases(c *gin.Context, patId string) ([]model.PatientDisease, error) {
	return s.d.GetPatientDiseases(c, patId)
}

func (s *patService) DeletePatientDisease(c *gin.Context, patId, diseaseId string) error {
	return s.d.DeletePatientDisease(c, patId, diseaseId)
}

func (s *patService) UpdatePatientDisease(c *gin.Context, patientDisease model.PatientDisease) error {
	return s.d.UpdatePatientDisease(c, patientDisease)
}

func (s *patService) RegisterPatientAllergy(c *gin.Context, p model.PatientAllergy) error {
	return s.d.RegisterPatientAllergy(c, p)
}

func (s *patService) GetPatientAllergies(c *gin.Context, patId string) ([]model.PatientAllergy, error) {
	return s.d.GetPatientAllergies(c, patId)
}

func (s *patService) DeletePatientAllergy(c *gin.Context, patId, allergyId string) error {
	return s.d.DeletePatientAllergy(c, patId, allergyId)
}

func (s *patService) UpdatePatientAllergy(c *gin.Context, patientAllergy model.PatientAllergy) error {
	return s.d.UpdatePatientAllergy(c, patientAllergy)
}

func (s *patService) RegisterPatientTreatment(c *gin.Context, p model.PatientTreatment) error {
	return s.d.RegisterPatientTreatment(c, p)
}

func (s *patService) GetPatientTreatments(c *gin.Context, patId string) ([]model.PatientTreatment, error) {
	return s.d.GetPatientTreatments(c, patId)
}

func (s *patService) DeletePatientTreatment(c *gin.Context, patId, treatmentId string) error {
	return s.d.DeletePatientTreatment(c, patId, treatmentId)
}

func (s *patService) UpdatePatientTreatment(c *gin.Context, patientTreatment model.PatientTreatment) error {
	return s.d.UpdatePatientTreatment(c, patientTreatment)
}

func (s *patService) RegisterPatientHistory(c *gin.Context, p model.PatientHistory) error {
	return s.d.RegisterPatientHistory(c, p)
}

func (s *patService) GetPatientHistory(c *gin.Context, patId string) ([]model.PatientHistory, error) {
	return s.d.GetPatientHistory(c, patId)
}

func (s *patService) DeletePatientHistory(c *gin.Context, patId, historyId string) error {
	return s.d.DeletePatientHistory(c, patId, historyId)
}

func (s *patService) UpdatePatientHistory(c *gin.Context, patientHistory model.PatientHistory) error {
	return s.d.UpdatePatientHistory(c, patientHistory)
}

func (s *patService) RegisterHospitalisation(c *gin.Context, p model.Hospitalisation) error {
	return s.d.RegisterHospitalisation(c, p)
}

func (s *patService) GetHospitalisations(c *gin.Context, patId string) ([]model.Hospitalisation, error) {
	return s.d.GetHospitalisations(c, patId)
}

func (s *patService) DeleteHospitalisation(c *gin.Context, patId, hospId string) error {
	return s.d.DeleteHospitalisation(c, patId, hospId)
}

func (s *patService) UpdateHospitalisation(c *gin.Context, hosp model.Hospitalisation) error {
	return s.d.UpdateHospitalisation(c, hosp)
}
