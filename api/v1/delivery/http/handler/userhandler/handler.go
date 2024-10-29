package userhandler

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
)

type Interactor interface {
	Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(ctx context.Context, req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(ctx context.Context, req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(ctx context.Context, req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
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
