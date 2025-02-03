package task

import (
	"context"
	"errors"

	taskdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/task"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Create(ctx context.Context, req taskdto.CreateRequest) (taskdto.CreateResponse, error) {
	isExistsUser, err := i.repository.IsExistsUser(ctx, req.UserID.Uint64())
	if err != nil {
		return taskdto.CreateResponse{}, err
	}

	if !isExistsUser {
		return taskdto.CreateResponse{}, richerror.New(_opTaskServiceCreate).
			WithErr(errors.New(_errMsgUserNotFound)).
			WithMessage(_errMsgUserNotFound).
			WithKind(richerror.KindStatusBadRequest)
	}

	t := entity.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.TaskPending,
	}

	createdTask, err := i.repository.Create(ctx, t)
	if err != nil {
		return taskdto.CreateResponse{}, err
	}

	return taskdto.CreateResponse{TaskInfo: taskdto.TaskInfo{
		ID:          createdTask.ID,
		UserID:      createdTask.UserID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}}, nil
}
