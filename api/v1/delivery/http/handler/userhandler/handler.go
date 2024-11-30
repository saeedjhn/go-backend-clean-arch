package userhandler

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract/tracercontract"

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

type Validator interface {
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error)
	ValidateCreateTaskRequest(req userdto.CreateTaskRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req userdto.RefreshTokenRequest) (map[string]string, error)
}

type Handler struct {
	app      *bootstrap.Application
	trc      tracercontract.Tracer
	vld      Validator
	userIntr Interactor
}

func New(
	app *bootstrap.Application,
	tracer tracercontract.Tracer,
	userValidator Validator,
	userInteractor Interactor,
) *Handler {
	return &Handler{
		app:      app,
		trc:      tracer,
		vld:      userValidator,
		userIntr: userInteractor,
	}
}
