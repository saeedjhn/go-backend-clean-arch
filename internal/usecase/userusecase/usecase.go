package userusecase

import (
	"go-backend-clean-arch/configs"
	"go-backend-clean-arch/internal/domain"
	"go-backend-clean-arch/internal/infrastructure/token"
)

type Repository interface {
	Create(u domain.User) (domain.User, error)
	IsMobileUnique(mobile string) (bool, error)
	GetByMobile(mobile string) (domain.User, error)
}

type Gateway interface {
	TaskList()
}

type UserInteractor struct {
	config          *configs.Config
	taskListGateway Gateway
	repository      Repository
	token           *token.Token
}

func New(
	config *configs.Config,
	taskListGateway Gateway,
	repository Repository,
	token *token.Token) *UserInteractor {
	return &UserInteractor{
		config:          config,
		taskListGateway: taskListGateway,
		repository:      repository,
		token:           token,
	}
}
