package admin

import (
	"context"

	admindto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/admin"
)

func (i *Interactor) Login(ctx context.Context, req admindto.LoginRequest) (admindto.LoginResponse, error) {
	panic("IMPLEMENT_ME")
}
