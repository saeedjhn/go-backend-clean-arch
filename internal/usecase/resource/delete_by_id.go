package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) DeleteByID(
	ctx context.Context,
	req resourcedto.DeleteByIDRequest,
) (resourcedto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
