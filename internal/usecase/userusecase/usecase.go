package userusecase

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
)

type AuthInteractor interface {
	CreateAccessToken(req entity.Authenticable) (string, error)
	CreateRefreshToken(req entity.Authenticable) (string, error)
	ParseToken(secret, requestToken string) (*authusecase.Claims, error)
}
type Repository interface {
	Create(ctx context.Context, u entity.User) (entity.User, error)
	IsMobileUnique(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

type Cache interface {
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) (bool, error)
}

type Interactor struct {
	config     *configs.Config
	authIntr   AuthInteractor
	cache      Cache
	repository Repository
}

// var _ userhandler.Interactor = (*Interactor)(nil) // Commented, because it happens import cycle.

func New(
	config *configs.Config,
	authIntr AuthInteractor,
	cache Cache,
	repository Repository,
) *Interactor {
	return &Interactor{
		config:     config,
		authIntr:   authIntr,
		cache:      cache,
		repository: repository,
	}
}
