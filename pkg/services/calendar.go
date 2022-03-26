package service

import (
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type CalService interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
}

type CalDb interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
}

type calService struct {
	d CalDb
}

func NewCalService(d CalDb) CalService {
	return &calService{d}
}

func (s *calService) GetAllEvents(c *gin.Context) ([]model.Event, error) {
	return s.d.GetAllEvents(c)
}
