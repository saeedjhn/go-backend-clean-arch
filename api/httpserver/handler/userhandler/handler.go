package userhandler

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
)

type Interactor interface {
	Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	TaskList(req userdto.TaskListRequest)
}

type UserHandler struct {
	app            *bootstrap.Application
	userInteractor Interactor
}

func New(app *bootstrap.Application, userInteractor Interactor) *UserHandler {
	return &UserHandler{app: app, userInteractor: userInteractor}
}
