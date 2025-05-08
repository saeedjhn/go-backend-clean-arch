package roleresourcepermission

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

func (i Interactor) GetByRoleIDsAndResourceID(
	_ context.Context,
	_ []uint64,
	_ uint64,
) (models.RoleResourcePermission, error) {
	panic("IMPLEMENT ME")
}
