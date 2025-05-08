package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/roleresourcepermission"
)

func (i Interactor) GetAll(
	_ context.Context,
	_ roleresourcepermissiondto.GetAllRequest,
) (roleresourcepermissiondto.GetAllResponse, error) {
	panic("IMPLEMENT ME")
}
