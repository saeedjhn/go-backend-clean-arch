package usergateway

import (
	"go-backend-clean-arch/internal/domain"
	"log"
)

type AuthInteractor interface {
	CreateAccessToken(user domain.User) (string, error)
	CreateRefreshToken(user domain.User) (string, error)
}

type TaskInteractor interface {
	List()
}

type UserGateway struct {
	authInteractor AuthInteractor
	taskInteractor TaskInteractor
}

func New(authInteractor AuthInteractor, taskInteractor TaskInteractor) *UserGateway {
	return &UserGateway{authInteractor: authInteractor, taskInteractor: taskInteractor}
}

func (u *UserGateway) TaskList() {
	u.taskInteractor.List()
	log.Print("TaskingGateway -> TaskList - IMPL ME")
}

func (u *UserGateway) CreateAccessToken(user domain.User) (string, error) {
	return u.authInteractor.CreateAccessToken(user)
}

func (u *UserGateway) CreateRefreshToken(user domain.User) (string, error) {
	return u.authInteractor.CreateRefreshToken(user)
}
