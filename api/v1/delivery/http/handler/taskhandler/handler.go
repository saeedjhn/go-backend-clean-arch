package taskhandler

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"
)

type Interactor interface {
	Create(
		ctx context.Context,
		dto usertaskservicedto.CreateTaskRequest,
	) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(
		ctx context.Context,
		dto usertaskservicedto.TasksUserRequest,
	) (usertaskservicedto.TasksUserResponse, error)
}

type Validator interface {
	ValidateCreateRequest(req taskdto.CreateRequest) (map[string]string, error)
}

type Handler struct {
	app      *bootstrap.Application
	vld      Validator
	taskIntr Interactor
}

func New(
	app *bootstrap.Application,
	validator Validator,
	taskInteractor Interactor,
) *Handler {
	return &Handler{
		app:      app,
		vld:      validator,
		taskIntr: taskInteractor,
	}
}
