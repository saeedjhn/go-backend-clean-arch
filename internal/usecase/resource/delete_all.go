package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) DeleteAll(
	ctx context.Context,
	req resourcedto.DeleteAllRequest,
) (resourcedto.DeleteAllResponse, error) {
	panic("IMPLEMENT ME")
}
