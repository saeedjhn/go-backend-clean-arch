package taskusecase

import (
	"context"
	"errors"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
)

func (i *Interactor) Create(
	ctx context.Context,
	req usertaskservicedto.CreateTaskRequest,
) (usertaskservicedto.CreateTaskResponse, error) {
	task := entity.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	isExistsUser, err := i.repository.IsExistsUser(req.UserID)
	if err != nil {
		return usertaskservicedto.CreateTaskResponse{}, err
	}

	if !isExistsUser {
		return usertaskservicedto.CreateTaskResponse{}, richerror.New(_opTaskServiceCreate).
			WithErr(errors.New(_errMsgUserNotFound)).
			WithMessage(_errMsgUserNotFound).
			WithKind(kind.KindStatusBadRequest)
	}

	createdTask, err := i.repository.Create(task)
	if err != nil {
		return usertaskservicedto.CreateTaskResponse{}, err
	}

	return usertaskservicedto.CreateTaskResponse{Task: createdTask}, nil
}
