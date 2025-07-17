package user

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
)

type Handler struct {
	trc      contract.Tracer
	userIntr userusecase.Interactor
}

func New(trc contract.Tracer, userIntr userusecase.Interactor) Handler {
	return Handler{trc: trc, userIntr: userIntr}
}
