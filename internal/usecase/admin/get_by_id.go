package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i Interactor) GetByID(
	_ context.Context,
	_ admindto.GetByIDRequest,
) (admindto.GetByIDResponse, error) {
	panic("IMPLEMENT_ME")
}
