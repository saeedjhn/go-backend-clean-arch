package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i Interactor) Register(
	_ context.Context,
	_ admindto.RegisterRequest,
) (admindto.RegisterResponse, error) {
	panic("IMPLEMENT_ME")
}
