package taskhandler

import (
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"go-backend-clean-arch/internal/domain/dto/taskdto"
)

type Interactor interface {
	Create(dto usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(dto usertaskservicedto.TasksUserRequest) (usertaskservicedto.TasksUserResponse, error)
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
