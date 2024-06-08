package userdto

import "go-backend-clean-arch/internal/domain"

type TasksRequest struct {
	ID uint `param:"id" json:"id"`
}

type TasksResponse struct {
	Tasks []domain.Task `json:"tasks"`
}
