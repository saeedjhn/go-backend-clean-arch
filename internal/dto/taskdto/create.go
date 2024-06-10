package taskdto

import "go-backend-clean-arch/internal/domain"

type CreateRequest struct {
	UserID      uint          `json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      domain.Status `json:"status"`
}

type CreateResponse struct {
	Task domain.Task `json:"task"`
}
