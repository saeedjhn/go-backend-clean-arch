package article

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
	"time"
)

type Tag struct {
	ID        types.ID
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
