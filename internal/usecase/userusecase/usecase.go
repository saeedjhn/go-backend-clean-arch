package userusecase

import (
	"go-backend-clean-arch/internal/domain"
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
	repository      Repository
	taskListGateway Gateway
}

func New(taskListGateway Gateway, repository Repository) *UserInteractor {
	return &UserInteractor{taskListGateway: taskListGateway, repository: repository}
}
