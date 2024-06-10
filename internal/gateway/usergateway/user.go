package usergateway

import (
	"go-backend-clean-arch/internal/contract"
	"go-backend-clean-arch/internal/domain"
)

type AuthContract interface {
	CreateAccessToken(user domain.User) (string, error)
	CreateRefreshToken(user domain.User) (string, error)
}

type TaskContract interface {
	TasksForUser(dto contract.GetTasksRequestDTO) (contract.GetTasksResponseDTO, error)
	Create(dto contract.CreateTaskRequestDTO) (contract.CreateTaskResponseDTO, error)
}

type UserGateway struct {
	auth AuthContract
	task TaskContract
}

func New(authInteractor AuthContract, taskInteractor TaskContract) *UserGateway {
	return &UserGateway{auth: authInteractor, task: taskInteractor}
}

func (u *UserGateway) Tasks(dto contract.GetTasksRequestDTO) (contract.GetTasksResponseDTO, error) {
	return u.task.TasksForUser(dto)
}

func (u *UserGateway) CreateTask(dto contract.CreateTaskRequestDTO) (contract.CreateTaskResponseDTO, error) {
	return u.task.Create(dto)
}

func (u *UserGateway) CreateAccessToken(user domain.User) (string, error) {
	return u.auth.CreateAccessToken(user)
}

func (u *UserGateway) CreateRefreshToken(user domain.User) (string, error) {
	return u.auth.CreateRefreshToken(user)
}
