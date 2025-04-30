package article

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
	"time"
)

type Category struct {
	ID          types.ID
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
