package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) GetByID(_ context.Context, _ resourcedto.GetByIDRequest) (resourcedto.GetByIDResponse, error) {
	panic("IMPLEMENT ME")
}
