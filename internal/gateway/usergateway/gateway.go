package usergateway

import (
	"log"
)

type TaskInteractor interface {
	List()
}

type Repository interface {
	Create()
}

type UserGateway struct {
	userRepository Repository
	taskInteractor TaskInteractor
}

func New(userRepository Repository, taskInteractor TaskInteractor) *UserGateway {
	return &UserGateway{userRepository: userRepository, taskInteractor: taskInteractor}
}

func (g UserGateway) SaveUserToDB() {
	g.userRepository.Create()

	// Any impl codes

	log.Print("UserGateway -> SaveUserToDB - IMPL ME")
}

func (g UserGateway) TaskList() {
	g.taskInteractor.List()

	// Any impl codes

	log.Print("UserGateway -> TaskList - IMPL ME")
}
