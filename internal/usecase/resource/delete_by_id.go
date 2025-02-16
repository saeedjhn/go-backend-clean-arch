package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ resourcedto.DeleteByIDRequest,
) (resourcedto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
