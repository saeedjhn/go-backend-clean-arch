package entity

import "time"

// RoleWildCardPermission defines permissions for a role on a pattern of resources.
type RoleWildCardPermission struct {
	RoleID          uint64      // Code of the role
	ResourcePattern string      // Resource pattern using wildcards (e.g., "resource.*")
	Permissions     Permissions // Allowed and denied actions
	Priority        int         // Priority of the wildcard permission
	CreatedAt       time.Time   // Timestamp for permission assignment
	UpdatedAt       time.Time   // Timestamp for the last permission update
}
