package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"
)

type AuthInteractor interface {
	CreateAccessToken(req entity.Authenticable) (string, error)
	CreateRefreshToken(req entity.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authusecase.Claims, error)
}

type UserInteractor interface {
	Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error)
	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
	Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error)
	RefreshToken(ctx context.Context, req user.RefreshTokenRequest) (user.RefreshTokenResponse, error)
}

type Handler struct {
	trc contract.Tracer
	// authIntr AuthInteractor
	// userIntr UserInteractor
	authIntr *authusecase.Interactor
	userIntr *userusecase.Interactor
}

func New(
	trc contract.Tracer,
	// authIntr AuthInteractor,
	// userIntr UserInteractor,
	authIntr *authusecase.Interactor,
	userIntr *userusecase.Interactor,
) *Handler {
	return &Handler{
		trc:      trc,
		authIntr: authIntr,
		userIntr: userIntr,
	}
}
