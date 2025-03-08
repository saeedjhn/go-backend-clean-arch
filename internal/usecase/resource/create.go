package resource

import (
	"context"

	resourcedto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/resource"
)

func (i *Interactor) Create(_ context.Context, _ resourcedto.CreateRequest) (resourcedto.CreateResponse, error) {
	panic("IMPLEMENT ME")
}
