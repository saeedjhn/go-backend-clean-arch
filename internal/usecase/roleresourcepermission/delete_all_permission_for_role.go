package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/roleresourcepermission"
)

func (i *Interactor) DeleteAllPermissionForRole(
	_ context.Context,
	_ roleresourcepermissiondto.DeleteAllPermissionForRoleRequest,
) (roleresourcepermissiondto.DeleteAllPermissionForRoleResponse, error) {
	panic("IMPLEMENT ME")
}
