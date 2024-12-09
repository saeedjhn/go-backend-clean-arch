package taskusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"
)

func (i *Interactor) FindAllByUserID(ctx context.Context, req taskdto.FindAllByUserIDRequest) (taskdto.FindAllByUserIDResponse, error) {
	tasks, err := i.repository.GetAllByUserID(ctx, req.UserID)
	if err != nil {
		return taskdto.FindAllByUserIDResponse{}, err
	}

	var data []taskdto.TaskInfo
	for _, task := range tasks {
		data = append(data, taskdto.TaskInfo{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			UpdatedAt:   task.UpdatedAt,
			CreatedAt:   task.CreatedAt,
		})
	}

	return taskdto.FindAllByUserIDResponse{Data: data}, nil
}
