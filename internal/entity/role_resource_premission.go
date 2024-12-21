package entity

import "time"

// RoleResourcePermission defines permissions for a role on a specific resource.
type RoleResourcePermission struct {
	RoleID      uint64      // Code of the role
	ResourceID  uint64      // Code of the resource
	Permissions Permissions // Allowed and denied actions
	CreatedAt   time.Time   // Timestamp for permission assignment
	UpdatedAt   time.Time   // Timestamp for the last permission update
}
