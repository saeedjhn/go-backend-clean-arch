package user

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
)

type UserInteractor interface {
	Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error)
	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
	Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error)
	RefreshToken(ctx context.Context, req user.RefreshTokenRequest) (user.RefreshTokenResponse, error)
}

type TaskInteractor interface {
	Create(ctx context.Context, req task.CreateRequest) (task.CreateResponse, error)
	FindAllByUserID(ctx context.Context, req task.FindAllByUserIDRequest) (task.FindAllByUserIDResponse, error)
}

type Validator interface {
	ValidateRegisterRequest(req user.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req user.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req user.ProfileRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req user.RefreshTokenRequest) (map[string]string, error)
	ValidateCreateTaskRequest(req task.CreateRequest) (map[string]string, error)
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
