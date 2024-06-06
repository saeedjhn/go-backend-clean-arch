package userhandler

import (
	"go-backend-clean-arch/internal/bootstrap"
	"go-backend-clean-arch/internal/dto/userdto"
)

type Interactor interface {
	Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(req userdto.LoginRequest) (userdto.LoginResponse, error)
	TaskList(req userdto.TaskListRequest)
}

type Validator interface {
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
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
