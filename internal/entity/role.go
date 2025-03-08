package entity

import (
	"time"

	"github.com/saeedjhn/go-domain-driven-design/internal/types"
)

// Role represents a stable set of permissions linked to business functions.
type Role struct {
	ID          types.ID // Unique code for the role
	Name        string   // Name displayed to end users
	Description string   // Overview of the role
	Internal    bool     // Indicates if the role is predefined and unmodifiable
	// GroupID     types.ID    // Group associated with this role
	CreatedAt time.Time // Timestamp for role creation
	UpdatedAt time.Time // Timestamp for the last role update
}
