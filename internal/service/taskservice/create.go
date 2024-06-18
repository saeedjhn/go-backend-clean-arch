package taskservice

import (
	"errors"
	"go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
	"go-backend-clean-arch/internal/domain/entity"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

func (t *TaskInteractor) Create(req usertaskservicedto.CreateTaskRequest) (usertaskservicedto.CreateTaskResponse, error) {
	const op = message.OpTaskUsecaseCreate

	task := entity.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	isExistsUser, err := t.repository.IsExistsUser(req.UserID)
	if err != nil {
		return usertaskservicedto.CreateTaskResponse{}, err
	}

	if !isExistsUser {
		return usertaskservicedto.CreateTaskResponse{}, richerror.New(op).
			WithErr(errors.New(message.ErrorMsgUserNotExists)).
			WithMessage(message.ErrorMsgInvalidInput).
			WithKind(kind.KindStatusBadRequest)
	}

	createdTask, err := t.repository.Create(task)
	if err != nil {
		return usertaskservicedto.CreateTaskResponse{}, err
	}

	return usertaskservicedto.CreateTaskResponse{Task: createdTask}, nil
}
