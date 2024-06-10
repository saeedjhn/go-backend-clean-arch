package userusecase

import (
	"go-backend-clean-arch/internal/contract"
	"go-backend-clean-arch/internal/dto/userdto"
)

func (u *UserInteractor) CreateTask(req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error) {
	dto := contract.CreateTaskRequestDTO{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	createdTask, err := u.gate.CreateTask(dto)
	if err != nil {
		return userdto.CreateTaskResponse{}, err
	}

	return userdto.CreateTaskResponse{Task: userdto.Task{
		ID:          createdTask.Task.ID,
		Title:       createdTask.Task.Title,
		Description: createdTask.Task.Description,
		Status:      createdTask.Task.Status,
		CreatedAt:   createdTask.Task.CreatedAt,
		UpdatedAt:   createdTask.Task.UpdatedAt,
	}}, nil
}
