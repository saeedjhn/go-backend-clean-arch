package admin

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
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
	Roles       []models.Role
	Gender      models.Gender
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
