package taskusecase

import (
	"errors"
	"go-backend-clean-arch/internal/contract"
	"go-backend-clean-arch/internal/domain"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

func (t *TaskInteractor) Create(dto contract.CreateTaskRequestDTO) (contract.CreateTaskResponseDTO, error) {
	const op = message.OpTaskUsecaseCreate

	createTask := domain.Task{
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
	}

	isExistsUser, err := t.repository.IsExistsUser(dto.UserID)
	if err != nil {
		return contract.CreateTaskResponseDTO{}, err
	}

	if !isExistsUser {
		return contract.CreateTaskResponseDTO{},
			richerror.New(op).
				WithErr(errors.New(message.ErrorMsgUserNotExists)). //nolint:err113
				WithMessage(message.ErrorMsgInvalidInput).
				WithKind(kind.KindStatusBadRequest)
	}

	createdTask, err := t.repository.Create(createTask)
	if err != nil {
		return contract.CreateTaskResponseDTO{}, err
	}

	return contract.CreateTaskResponseDTO{Task: contract.Task{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}}, nil
}
