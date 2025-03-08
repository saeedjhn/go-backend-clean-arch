package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-domain-driven-design/internal/dto/roleresourcepermission"
)

func (i *Interactor) DeleteAllPermissionForRole(
	_ context.Context,
	_ roleresourcepermissiondto.DeleteAllPermissionForRoleRequest,
) (roleresourcepermissiondto.DeleteAllPermissionForRoleResponse, error) {
	panic("IMPLEMENT ME")
}
