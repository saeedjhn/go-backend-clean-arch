package roleresourcepermission

import (
	"context"

	roleresourcepermissiondto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/roleresourcepermission"
)

func (i *Interactor) GetByID(
	ctx context.Context,
	req roleresourcepermissiondto.GetByIDRequest,
) (roleresourcepermissiondto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
