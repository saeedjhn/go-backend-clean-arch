package roleresourcepermission

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

func (i *Interactor) GetByRoleIDAndResourceID(
	ctx context.Context,
	roleID,
	resourceId uint64,
) (entity.RoleResourcePermission, error) {
	panic("IMPLEMENT ME")
}
