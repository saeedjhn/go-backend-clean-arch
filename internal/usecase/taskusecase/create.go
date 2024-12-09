package taskusecase

import (
	"context"
	"errors"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/taskdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) Create(ctx context.Context, req taskdto.CreateRequest) (taskdto.CreateResponse, error) {
	task := entity.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.TaskPending,
	}

	isExistsUser, err := i.repository.IsExistsUser(ctx, req.UserID)
	if err != nil {
		return taskdto.CreateResponse{}, err
	}

	if !isExistsUser {
		return taskdto.CreateResponse{}, richerror.New(_opTaskServiceCreate).
			WithErr(errors.New(_errMsgUserNotFound)).
			WithMessage(_errMsgUserNotFound).
			WithKind(kind.KindStatusBadRequest)
	}

	createdTask, err := i.repository.Create(ctx, task)
	if err != nil {
		return taskdto.CreateResponse{}, err
	}

	return taskdto.CreateResponse{Data: taskdto.TaskInfo{
		ID:          createdTask.ID,
		UserID:      createdTask.UserID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}}, nil
}
