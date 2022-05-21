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
