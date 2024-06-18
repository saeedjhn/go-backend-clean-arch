package userhandler

import (
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/domain/dto/userdto"
)

type Interactor interface {
	Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
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
	userValidator  Validator
	userInteractor Interactor
}

func New(
	app *bootstrap.Application,
	userValidator Validator,
	userInteractor Interactor,
) *UserHandler {
	return &UserHandler{
		app:            app,
		userValidator:  userValidator,
		userInteractor: userInteractor,
	}
}
