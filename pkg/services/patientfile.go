package patientfile

import (
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
}
type Db interface {
	GetAllPatients(c *gin.Context) ([]model.Patient, error)
	CreatePatient(c *gin.Context, p model.Patient) error
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
