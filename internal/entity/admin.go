package entity

import "time"

// Admin represents a system administrator with specific roles and permissions.
type Admin struct {
	ID          uint64
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
