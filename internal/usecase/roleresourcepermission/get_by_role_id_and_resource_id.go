package roleresourcepermission

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
)

func (i *Interactor) GetByRoleIDAndResourceID(
	_ context.Context,
	_ uint64,
	_ uint64,
) (entity.RoleResourcePermission, error) {
	panic("IMPLEMENT ME")
}
