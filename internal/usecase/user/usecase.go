package user

import (
	"context"
	entity2 "github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/auth"
)

type AuthInteractor interface {
	CreateAccessToken(req entity2.Authenticable) (string, error)
	CreateRefreshToken(req entity2.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*auth.Claims, error)
}
type Repository interface {
	Create(ctx context.Context, u entity2.User) (entity2.User, error)
	IsMobileUnique(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (entity2.User, error)
	GetByID(ctx context.Context, id uint64) (entity2.User, error)
}

type Cache interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (bool, error)
}

type Interactor struct {
	cfg        *configs.Config
	trc        contract.Tracer
	authIntr   AuthInteractor
	cache      Cache
	repository Repository
}

// var _ userhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	cfg *configs.Config,
	trc contract.Tracer,
	authIntr AuthInteractor,
	cache Cache,
	repository Repository,
) *Interactor {
	return &Interactor{
		cfg:        cfg,
		trc:        trc,
		authIntr:   authIntr,
		cache:      cache,
		repository: repository,
	}
}
