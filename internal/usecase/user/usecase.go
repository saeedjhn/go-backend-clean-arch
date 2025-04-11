package user

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	usermodel "github.com/saeedjhn/go-backend-clean-arch/internal/models/user"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authentication"
)

//go:generate mockery --name AuthInteractor
type AuthInteractor interface {
	CreateAccessToken(req models.Authenticable) (string, error)
	CreateRefreshToken(req models.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authentication.Claims, error)
}

//go:generate mockery --name Validator
type Validator interface {
	ValidateRegisterRequest(req userdto.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req userdto.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req userdto.ProfileRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req userdto.RefreshTokenRequest) (map[string]string, error)
}

//go:generate mockery --name Cache
type Cache interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (bool, error)
}

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, u usermodel.User) (usermodel.User, error)
	IsExistsByMobile(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (usermodel.User, error)
	GetByID(ctx context.Context, id uint64) (usermodel.User, error)
}

type Interactor struct {
	cfg        *configs.Config
	trc        contract.Tracer
	authIntr   AuthInteractor
	vld        Validator
	cache      Cache
	repository Repository
}

// var _ userhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	cfg *configs.Config,
	trc contract.Tracer,
	authIntr AuthInteractor,
	vld Validator,
	cache Cache,
	repository Repository,
) *Interactor {
	return &Interactor{
		cfg:        cfg,
		trc:        trc,
		vld:        vld,
		authIntr:   authIntr,
		cache:      cache,
		repository: repository,
	}
}

func GenerateHash(password string) (string, error) {
	return bcrypt.Generate(password, bcrypt.Cost(configs.BcryptCost))
}

func CompareHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndSTR(hashedPassword, password)
}
