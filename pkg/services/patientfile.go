package patientfile

import (
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string) ([]model.Patient, error)
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error)
}
type Db interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByName(c *gin.Context, nameOrId string) ([]model.Patient, error)
	GetPatientById(c *gin.Context, id string) (model.Patient, error)
	UpdatePatient(c *gin.Context, p model.Patient) error
	SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error)
}

type service struct {
	d Db
}

func NewService(d Db) Service {
	return &service{d}
}

func (s *service) GetAllPatients(c *gin.Context) ([]model.Patient, error) {
	return s.d.GetAllPatients(c)
}

func (s *service) CreatePatient(c *gin.Context, p model.Patient) error {
	return s.d.CreatePatient(c, p)
}

func (s *service) SearchPatientByName(c *gin.Context, nameOrId string) ([]model.Patient, error) {
	return s.d.SearchPatientByName(c, nameOrId)
}

func (s *service) GetPatientById(c *gin.Context, id string) (model.Patient, error) {
	return s.d.GetPatientById(c, id)
}

func (s *service) UpdatePatient(c *gin.Context, p model.Patient) error {
	return s.d.UpdatePatient(c, p)
}

func (s *service) SearchPatientByINSMatricule(c *gin.Context, id string) ([]model.Patient, error) {
	return s.d.SearchPatientByINSMatricule(c, id)
}
