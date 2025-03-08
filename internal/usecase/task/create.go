package task

import (
	"context"
	"errors"

	taskdto "github.com/saeedjhn/go-domain-driven-design/internal/dto/task"
	"github.com/saeedjhn/go-domain-driven-design/internal/entity"

	"github.com/saeedjhn/go-domain-driven-design/pkg/richerror"
)

func (i *Interactor) Create(ctx context.Context, req taskdto.CreateRequest) (taskdto.CreateResponse, error) {
	isExistsUser, err := i.repository.IsExistsUser(ctx, req.UserID.Uint64())
	if err != nil {
		return taskdto.CreateResponse{}, err
	}

	if !isExistsUser {
		return taskdto.CreateResponse{}, richerror.New(_opTaskServiceCreate).
			WithErr(errors.New(errMsgUserNotFound)).
			WithMessage(errMsgUserNotFound).
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

	return taskdto.CreateResponse{TaskInfo: taskdto.Info{
		ID:          createdTask.ID,
		UserID:      createdTask.UserID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}}, nil
}
