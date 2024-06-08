package usergateway

import (
	"go-backend-clean-arch/internal/domain"
)

type AuthContract interface {
	CreateAccessToken(user domain.User) (string, error)
	CreateRefreshToken(user domain.User) (string, error)
}

type TaskContract interface {
	TasksForUser(userID uint) ([]domain.Task, error)
}

type UserGateway struct {
	auth AuthContract
	task TaskContract
}

func New(authInteractor AuthContract, taskInteractor TaskContract) *UserGateway {
	return &UserGateway{auth: authInteractor, task: taskInteractor}
}

func (u *UserGateway) Tasks(userID uint) ([]domain.Task, error) {
	return u.task.TasksForUser(userID)
}

func (u *UserGateway) CreateAccessToken(user domain.User) (string, error) {
	return u.auth.CreateAccessToken(user)
}

func (u *UserGateway) CreateRefreshToken(user domain.User) (string, error) {
	return u.auth.CreateRefreshToken(user)
}
