package task

import (
	"context"

	task "github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
)

func (i *Interactor) FindAllByUserID(ctx context.Context, req task.FindAllByUserIDRequest) (task.FindAllByUserIDResponse, error) {
	tasks, err := i.repository.GetAllByUserID(ctx, req.UserID)
	if err != nil {
		return task.FindAllByUserIDResponse{}, err
	}

	var data []task.Data
	for _, t := range tasks {
		data = append(data, task.Data{
			ID:          t.ID,
			UserID:      t.UserID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			UpdatedAt:   t.UpdatedAt,
			CreatedAt:   t.CreatedAt,
		})
	}

	return task.FindAllByUserIDResponse{Data: data}, nil
}
