package taskusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
)

func (i *Interactor) TasksUser(
	ctx context.Context,
	req usertaskservicedto.TasksUserRequest,
) (usertaskservicedto.TasksUserResponse, error) {
	tasks, err := i.repository.GetAllByUserID(ctx, req.UserID)
	if err != nil {
		return usertaskservicedto.TasksUserResponse{}, err
	}

	return usertaskservicedto.TasksUserResponse{Tasks: tasks}, nil
}
