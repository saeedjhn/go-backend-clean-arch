package entity

import "time"

// Resource represents a protectable entity within the system.
type Resource struct {
	ID          uint64    // Unique code for the resource
	Name        string    // Name displayed to end users
	Description string    // Overview of the resource
	Type        string    // Type of resource (e.g., "module", "file", "API")
	Internal    bool      // Indicates if the resource is predefined and unmodifiable
	CreatedAt   time.Time // Timestamp for resource creation
	UpdatedAt   time.Time // Timestamp for the last resource update
}
