package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i *Interactor) DeleteByID(
	_ context.Context,
	_ admindto.DeleteByIDRequest,
) (admindto.DeleteByIDResponse, error) {
	panic("IMPLEMENT ME")
}
