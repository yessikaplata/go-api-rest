package handler

import (
	"context"

	"github.com/yessikaplata/go-api-rest/model"
	"github.com/yessikaplata/go-api-rest/service"
)

// handler
type Handler struct {
	service Service
}

// service
type Service interface {
	Create(ctx context.Context, params service.CreateParams) (model.User, error)
	Get(ctx context.Context, id int) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, params service.UpdateParams) (model.User, error)
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}
