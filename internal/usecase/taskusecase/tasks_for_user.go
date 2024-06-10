package taskusecase

import (
	"go-backend-clean-arch/internal/contract"
)

func (t *TaskInteractor) TasksForUser(dto contract.GetTasksRequestDTO) (contract.GetTasksResponseDTO, error) {
	tasks, err := t.repository.GetAllByUserID(dto.UserID)
	if err != nil {
		return contract.GetTasksResponseDTO{}, err
	}

	var taskRecord []contract.Task
	for _, task := range tasks {
		taskRecord = append(taskRecord, contract.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return contract.GetTasksResponseDTO{Tasks: taskRecord}, err
}
