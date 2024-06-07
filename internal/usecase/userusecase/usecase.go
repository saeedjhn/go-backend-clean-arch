package userusecase

import (
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/domain"
)

type Repository interface {
	Create(u domain.User) (domain.User, error)
	IsMobileUnique(mobile string) (bool, error)
	GetByMobile(mobile string) (domain.User, error)
	GetByID(id uint) (domain.User, error)
}

type Gateway interface {
	TaskList()
	CreateAccessToken(user domain.User) (string, error)
	CreateRefreshToken(user domain.User) (string, error)
}

type UserInteractor struct {
	config     *configs.Config
	gate       Gateway
	repository Repository
}

func New(
	config *configs.Config,
	gate Gateway,
	repository Repository,
) *UserInteractor {
	return &UserInteractor{
		config:     config,
		gate:       gate,
		repository: repository,
	}
}
