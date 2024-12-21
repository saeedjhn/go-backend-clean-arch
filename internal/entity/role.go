package entity

import "time"

// Role represents a stable set of permissions linked to business functions.
type Role struct {
	ID          uint64    // Unique code for the role
	Name        string    // Name displayed to end users
	Description string    // Overview of the role
	Internal    bool      // Indicates if the role is predefined and unmodifiable
	GroupID     uint64    // Group associated with this role
	CreatedAt   time.Time // Timestamp for role creation
	UpdatedAt   time.Time // Timestamp for the last role update
}
