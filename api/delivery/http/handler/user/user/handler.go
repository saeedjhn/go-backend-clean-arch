package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	authusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
)

// type AuthInteractor interface {
// 	CreateAccessToken(req models.Authenticable) (string, error)
// 	CreateRefreshToken(req models.Authenticable) (string, error)
// 	ParseToken(secret, requestToken string) (*authusecase.Claims, error)
// }
//
// type Interactor interface {
// 	Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error)
// 	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
// 	Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error)
// 	RefreshToken(ctx context.Context, req user.RefreshTokenRequest) (user.RefreshTokenResponse, error)
// }

type Handler struct {
	trc      contract.Tracer
	authIntr authusecase.Interactor
	userIntr userusecase.Interactor
}

func New(
	trc contract.Tracer,
	authIntr authusecase.Interactor,
	userIntr userusecase.Interactor,
) Handler {
	return Handler{
		trc:      trc,
		authIntr: authIntr,
		userIntr: userIntr,
	}
}
