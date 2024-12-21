package entity

import "time"

// Group represents a collection of roles for simplifying permission management.
type Group struct {
	ID          uint64    // Unique code for the group
	Name        string    // Name displayed to end users
	Description string    // Overview of the group
	Internal    bool      // Indicates if the group is predefined and unmodifiable
	Owner       string    // Identifier of the group owner
	CreatedAt   time.Time // Timestamp for group creation
	UpdatedAt   time.Time // Timestamp for the last group update
}
