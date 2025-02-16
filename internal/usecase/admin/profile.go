package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i *Interactor) Profile(
	_ context.Context,
	_ admindto.ProfileRequest,
) (admindto.ProfileResponse, error) {
	panic("IMPLEMENT_ME")
}
