package entity

import (
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/types"
)

// RoleResourcePermission defines permissions for a role on a specific resource.
type RoleResourcePermission struct {
	RoleID      types.ID   // Code of the role
	ResourceID  types.ID   // Code of the resource
	Permissions Permission // Allowed and denied actions
	CreatedAt   time.Time  // Timestamp for permission assignment
	UpdatedAt   time.Time  // Timestamp for the last permission update
}
