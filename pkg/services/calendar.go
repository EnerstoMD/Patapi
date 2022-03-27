package service

import (
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type CalService interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
}

type CalDb interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
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

func (s *calService) CreateEvent(c *gin.Context, e model.Event) error {
	return s.d.CreateEvent(c, e)
}

func (s *calService) UpdateEvent(c *gin.Context, e model.Event) error {
	return s.d.UpdateEvent(c, e)
}
