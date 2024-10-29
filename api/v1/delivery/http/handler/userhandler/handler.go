package userhandler

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
)

type Interactor interface {
	Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
}

type Presenter interface {
	Success(data interface{}) map[string]interface{}
	SuccessWithMSG(msg string, data interface{}) map[string]interface{}
	Error(err error) (int, map[string]interface{})
	ErrorWithMSG(msg string, err error) map[string]interface{}
}

type Validator interface {
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error)
	ValidateCreateTaskRequest(req userdto.CreateTaskRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req userdto.RefreshTokenRequest) (map[string]string, error)
}

type UserHandler struct {
	app            *bootstrap.Application
	present        Presenter
	userValidator  Validator
	userInteractor Interactor
}

func New(
	app *bootstrap.Application,
	present Presenter,
	userValidator Validator,
	userInteractor Interactor,
) *UserHandler {
	return &UserHandler{
		app:            app,
		present:        present,
		userValidator:  userValidator,
		userInteractor: userInteractor,
	}
}
