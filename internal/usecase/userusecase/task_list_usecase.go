package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"log"
)

func (u *UserInteractor) TaskList(req userdto.UserRequest) {
	u.userGateway.TaskList()

	log.Print("UserInteractor -> TaskList - IMPL ME")
}
