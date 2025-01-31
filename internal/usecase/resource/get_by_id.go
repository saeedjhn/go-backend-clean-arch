package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) GetByID(ctx context.Context, req resourcedto.GetByIDRequest) (resourcedto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
