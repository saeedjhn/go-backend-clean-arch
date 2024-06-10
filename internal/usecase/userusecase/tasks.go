package userusecase

import (
	"go-backend-clean-arch/internal/contract"
	"go-backend-clean-arch/internal/dto/userdto"
)

func (u *UserInteractor) Tasks(req userdto.TasksRequest) (userdto.TasksResponse, error) {
	dto := contract.GetTasksRequestDTO{
		UserID: req.ID,
	}

	tasks, err := u.gate.Tasks(dto)
	if err != nil {
		return userdto.TasksResponse{}, err
	}

	var taskRecord []userdto.Task
	for _, task := range tasks.Tasks {
		taskRecord = append(taskRecord, userdto.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return userdto.TasksResponse{Tasks: taskRecord}, nil
}
