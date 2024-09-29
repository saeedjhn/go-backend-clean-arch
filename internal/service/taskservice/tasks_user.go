package taskservice

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/usertaskservicedto"
)

func (t *TaskInteractor) TasksUser(
	req usertaskservicedto.TasksUserRequest,
) (usertaskservicedto.TasksUserResponse, error) {
	tasks, err := t.repository.GetAllByUserID(req.UserID)
	if err != nil {
		return usertaskservicedto.TasksUserResponse{}, err
	}

	return usertaskservicedto.TasksUserResponse{Tasks: tasks}, nil
}
