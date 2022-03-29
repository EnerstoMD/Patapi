package service

import (
	"lupus/patapi/pkg/model"

	"github.com/gin-gonic/gin"
)

type CalService interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
	DeleteEvent(ctx *gin.Context, id string) error
	ConfirmEvent(ctx *gin.Context, id string) error
	UnconfirmEvent(ctx *gin.Context, id string) error
}

type CalDb interface {
	GetAllEvents(c *gin.Context) ([]model.Event, error)
	CreateEvent(c *gin.Context, e model.Event) error
	UpdateEvent(c *gin.Context, e model.Event) error
	DeleteEvent(ctx *gin.Context, id string) error
	ConfirmEvent(ctx *gin.Context, id string) error
	UnconfirmEvent(ctx *gin.Context, id string) error
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

func (s *calService) DeleteEvent(ctx *gin.Context, id string) error {
	return s.d.DeleteEvent(ctx, id)
}

func (s *calService) ConfirmEvent(ctx *gin.Context, id string) error {
	return s.d.ConfirmEvent(ctx, id)
}

func (s *calService) UnconfirmEvent(ctx *gin.Context, id string) error {
	return s.d.UnconfirmEvent(ctx, id)
}
