package models

import (
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

// Resource represents a protectable models within the system.
type Resource struct {
	ID          types.ID
	Name        string
	Description string
	Type        string // Type of resource (e.g., "module", "file", "API")
	Internal    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ResourceStatus string

const (
	ResourceModuleType = ResourceStatus("module")
	ResourceAPIType    = ResourceStatus("API")
	ResourceFileType   = ResourceStatus("file")
)

var _resourceTypeStrings = map[ResourceStatus]string{ //nolint:gochecknoglobals // nothing
	ResourceModuleType: "module",
	ResourceAPIType:    "API",
	ResourceFileType:   "file",
}

func (a ResourceStatus) IsValid() bool {
	_, ok := _resourceTypeStrings[a]

	return ok
}
