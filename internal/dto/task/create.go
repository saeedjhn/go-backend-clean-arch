package task

import "github.com/saeedjhn/go-domain-driven-design/internal/types"

type CreateRequest struct {
	UserID      types.ID `json:"user_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	// Status      uint8  `json:"status"`
}

type CreateResponse struct {
	TaskInfo    Info              `json:"task"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
