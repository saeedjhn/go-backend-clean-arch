package user

import (
	"context"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
	authusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/authentication"
	userusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/user"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto/user"
)

type AuthInteractor interface {
	CreateAccessToken(req entity.Authenticable) (string, error)
	CreateRefreshToken(req entity.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authusecase.Claims, error)
}

type Interactor interface {
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
