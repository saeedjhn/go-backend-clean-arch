package user

import (
	"context"
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-domain-driven-design/pkg/security/bcrypt"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto/user"

	"github.com/saeedjhn/go-domain-driven-design/internal/entity"

	"github.com/saeedjhn/go-domain-driven-design/configs"
	"github.com/saeedjhn/go-domain-driven-design/internal/usecase/authentication"
)

//go:generate mockery --name AuthInteractor
type AuthInteractor interface {
	CreateAccessToken(req entity.Authenticable) (string, error)
	CreateRefreshToken(req entity.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authentication.Claims, error)
}

//go:generate mockery --name Validator
type Validator interface {
	ValidateRegisterRequest(req user.RegisterRequest) (map[string]string, error)
	ValidateLoginRequest(req user.LoginRequest) (map[string]string, error)
	ValidateProfileRequest(req user.ProfileRequest) (map[string]string, error)
	ValidateRefreshTokenRequest(req user.RefreshTokenRequest) (map[string]string, error)
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
	Create(ctx context.Context, u entity.User) (entity.User, error)
	IsExistsByMobile(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
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
