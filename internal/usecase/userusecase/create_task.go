package userusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
)

func (i *Interactor) CreateTask(ctx context.Context, req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error) {
	dto := usertaskservicedto.CreateTaskRequest{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.Pending,
	}

	createdTask, err := i.taskIntr.Create(ctx, dto)
	if err != nil {
		return userdto.CreateTaskResponse{}, err
	}

	return userdto.CreateTaskResponse{Task: createdTask.Task}, nil
}
