package article

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Tag struct {
	ID        types.ID
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
