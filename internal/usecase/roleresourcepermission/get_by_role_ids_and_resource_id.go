package roleresourcepermission

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

func (i *Interactor) GetByRoleIDsAndResourceID(
	ctx context.Context,
	roleIDs []uint64,
	resourceId uint64,
) (entity.RoleResourcePermission, error) {
	panic("IMPLEMENT ME")
}
