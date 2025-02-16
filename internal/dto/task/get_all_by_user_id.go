package task

import "github.com/saeedjhn/go-backend-clean-arch/internal/types"

type GetAllByUserIDRequest struct {
	UserID types.ID `param:"id" json:"user_id"`
}

type GetByUserIDResponse struct {
	Tasks       []Info            `json:"tasks"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
