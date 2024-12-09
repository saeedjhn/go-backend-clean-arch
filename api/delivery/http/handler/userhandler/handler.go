package userhandler

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
)

type UserInteractor interface {
	Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(ctx context.Context, req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
}

type TaskInteractor interface {
	Create(ctx context.Context, req taskdto.CreateRequest) (taskdto.CreateResponse, error)
	FindAllByUserID(ctx context.Context, req taskdto.FindAllByUserIDRequest) (taskdto.FindAllByUserIDResponse, error)
}

type Validator interface {
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req userdto.RefreshTokenRequest) (map[string]string, error)
	ValidateCreateTaskRequest(req taskdto.CreateRequest) (map[string]string, error)
}

type Handler struct {
	app      *bootstrap.Application
	trc      contract.Tracer
	vld      Validator
	userIntr UserInteractor
	taskIntr TaskInteractor
}

func New(
	app *bootstrap.Application,
	trc contract.Tracer,
	userValidator Validator,
	userInteractor UserInteractor,
	taskInteractor TaskInteractor,
) *Handler {
	return &Handler{
		app:      app,
		trc:      trc,
		vld:      userValidator,
		userIntr: userInteractor,
		taskIntr: taskInteractor,
	}
}
