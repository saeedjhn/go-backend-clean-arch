package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/roleresourcepermission"
)

func (i Interactor) DeleteAll(
	_ context.Context,
	_ roleresourcepermissiondto.DeleteAllRequest,
) (roleresourcepermissiondto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
