package task

import (
	"context"

	taskdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
)

func (i *Interactor) GetAllByUserID(ctx context.Context, req taskdto.GetAllByUserIDRequest) (taskdto.GetByUserIDResponse, error) {
	tasksByUserID, err := i.repository.GetAllByUserID(ctx, req.UserID.Uint64())
	if err != nil {
		return taskdto.GetByUserIDResponse{}, err
	}

	var tasks []taskdto.TaskInfo
	for _, task := range tasksByUserID {
		tasks = append(tasks, taskdto.TaskInfo{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			UpdatedAt:   task.UpdatedAt,
			CreatedAt:   task.CreatedAt,
		})
	}

	return taskdto.GetByUserIDResponse{Tasks: tasks}, nil
}
