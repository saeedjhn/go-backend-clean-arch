package userusecase

import (
	"go-backend-clean-arch/internal/dto/userdto"
	"log"
)

func (u *UserInteractor) TaskList(req userdto.TaskListRequest) {
	u.taskListGateway.TaskList()

	log.Print("UserInteractor -> TaskList - IMPL ME")
}
