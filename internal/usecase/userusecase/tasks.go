package userusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
)

func (i *Interactor) Tasks(ctx context.Context, req userdto.TasksRequest) (userdto.TasksResponse, error) {
	dto := usertaskservicedto.TasksUserRequest{UserID: req.ID}

	tasksUser, err := i.taskIntr.TasksUser(ctx, dto)
	if err != nil {
		return userdto.TasksResponse{}, err
	}

	var resp []userdto.TaskInfo
	for _, task := range tasksUser.Tasks {
		resp = append(resp, userdto.TaskInfo{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return userdto.TasksResponse{Tasks: resp}, nil
}
