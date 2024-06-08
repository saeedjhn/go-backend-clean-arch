package userusecase

import (
	"go-backend-clean-arch/internal/dto/userdto"
)

func (u *UserInteractor) Tasks(req userdto.TasksRequest) (userdto.TasksResponse, error) {
	tasks, _ := u.gate.Tasks(req.ID) // escape err

	return userdto.TasksResponse{Tasks: tasks}, nil
}
