package task

import (
	"context"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Create(ctx context.Context, req task.CreateRequest) (task.CreateResponse, error) {
	isExistsUser, err := i.repository.IsExistsUser(ctx, req.UserID)
	if err != nil {
		return task.CreateResponse{}, err
	}

	if !isExistsUser {
		return task.CreateResponse{}, richerror.New(_opTaskServiceCreate).
			WithErr(errors.New(_errMsgUserNotFound)).
			WithMessage(_errMsgUserNotFound).
			WithKind(kind.KindStatusBadRequest)
	}

	t := entity.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.TaskPending,
	}

	createdTask, err := i.repository.Create(ctx, t)
	if err != nil {
		return task.CreateResponse{}, err
	}

	return task.CreateResponse{Data: task.Data{
		ID:          createdTask.ID,
		UserID:      createdTask.UserID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}}, nil
}
