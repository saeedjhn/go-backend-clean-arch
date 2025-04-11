package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type ProfileRequest struct {
	ID types.ID `json:"id"`
}

type ProfileResponse struct {
	UserInfo    Info              `json:"user"`
	FieldErrors map[string]string `json:"field_errors,omitempty"`
}
