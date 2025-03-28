package entity

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"
)

// Admin represents a system administrator with specific roles and permissions.
type Admin struct {
	ID          types.ID
	FirstName   string
	LastName    string
	Email       string
	Mobile      string
	Description string
	Password    string
	Roles       []Role
	Gender      Gender
	Status      AdminStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AdminStatus string

const (
	AdminActiveStatus   = AdminStatus("active")
	AdminInactiveStatus = AdminStatus("inactive")
)

var _adminStatusStrings = map[AdminStatus]string{ //nolint:gochecknoglobals // nothing
	AdminActiveStatus:   "active",
	AdminInactiveStatus: "inactive",
}

func (a AdminStatus) IsValid() bool {
	_, ok := _adminStatusStrings[a]

	return ok
}
