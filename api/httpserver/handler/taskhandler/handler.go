package taskhandler

import (
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/dto/taskdto"
)

type Interactor interface {
	Create(req taskdto.CreateRequest) (taskdto.CreateResponse, error)
}

type Validator interface {
	ValidateCreateRequest(req taskdto.CreateRequest) (map[string]string, error)
}

type TaskHandler struct {
	app            *bootstrap.Application
	validator      Validator
	taskInteractor Interactor
}

func New(
	app *bootstrap.Application,
	validator Validator,
	taskInteractor Interactor,
) *TaskHandler {
	return &TaskHandler{
		app:            app,
		validator:      validator,
		taskInteractor: taskInteractor,
	}
}
