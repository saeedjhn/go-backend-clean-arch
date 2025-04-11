package models

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

// RoleWildCardPermission defines permissions for a role on a pattern of resources.
type RoleWildCardPermission struct {
	RoleID          types.ID   // Code of the role
	ResourcePattern string     // Resource pattern using wildcards (e.g., "resource.*")
	Priority        int        // Priority of the wildcard permission
	Permissions     Permission // Allowed and denied actions
	CreatedAt       time.Time  // Timestamp for permission assignment
	UpdatedAt       time.Time  // Timestamp for the last permission update
}
