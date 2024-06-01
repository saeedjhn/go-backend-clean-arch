package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
)

type Repository interface {
	Register(u domain.User) (domain.User, error)
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
