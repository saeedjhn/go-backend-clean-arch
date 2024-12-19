package task

import (
	"context"
	task2 "github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
)

type Interactor interface {
	Create(ctx context.Context, req task2.CreateRequest) (task2.CreateResponse, error)
	FindAll(ctx context.Context, req task2.FindAllRequest) (task2.FindAllResponse, error)
	FindAllByUserID(ctx context.Context, req task2.FindAllByUserIDRequest) (task2.FindAllByUserIDResponse, error)
}

type Validator interface {
	ValidateCreateRequest(req task2.CreateRequest) (map[string]string, error)
}

type Handler struct {
	app      *bootstrap.Application
	trc      contract.Tracer
	vld      Validator
	taskIntr Interactor
}

func New(
	app *bootstrap.Application,
	trc contract.Tracer,
	validator Validator,
	taskInteractor Interactor,
) *Handler {
	return &Handler{
		app:      app,
		trc:      trc,
		vld:      validator,
		taskIntr: taskInteractor,
	}
}
