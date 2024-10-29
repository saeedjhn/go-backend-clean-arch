package taskhandler

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"
)

type Interactor interface {
	Create(dto usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error)
	TasksUser(dto usertaskservicedto.TasksUserRequest) (usertaskservicedto.TasksUserResponse, error)
}

type Presenter interface {
	Success(data interface{}) map[string]interface{}
	SuccessWithMSG(msg string, data interface{}) map[string]interface{}
	Error(err error) (int, map[string]interface{})
	ErrorWithMSG(msg string, err error) map[string]interface{}
}

type Validator interface {
	ValidateCreateRequest(req taskdto.CreateRequest) (map[string]string, error)
}

type TaskHandler struct {
	app            *bootstrap.Application
	present        Presenter
	validator      Validator
	taskInteractor Interactor
}

func New(
	app *bootstrap.Application,
	present Presenter,
	validator Validator,
	taskInteractor Interactor,
) *TaskHandler {
	return &TaskHandler{
		app:            app,
		present:        present,
		validator:      validator,
		taskInteractor: taskInteractor,
	}
}
