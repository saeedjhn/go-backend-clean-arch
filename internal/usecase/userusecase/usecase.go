package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"golang.org/x/crypto/bcrypt"
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

func bcryptPassword(password string) string {
	encryptedPass, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // TODO - implement dynamic cost
	)

	return string(encryptedPass)
}
