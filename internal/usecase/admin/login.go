package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-domain-driven-design/internal/dto/admin"
)

func (i *Interactor) Login(
	_ context.Context,
	_ admindto.LoginRequest,
) (admindto.LoginResponse, error) {
	panic("IMPLEMENT_ME")
}
