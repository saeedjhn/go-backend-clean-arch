package user

import (
	"context"
	"log"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
)

func (i Interactor) RegisteredHandler(ctx context.Context, req events.UserRegisteredEvent) error {
	_, span := i.trc.Span(ctx, "RegisteredHandler")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "RegisteredHander",
		"usecase.request": req,
	})
	defer span.End()

	// if fieldsErrs, err := i.vld.ValidateRegisterRequest(req); err != nil {
	// 	return userdto.RegisterResponse{FieldErrors: fieldsErrs}, err
	// }

	log.Printf("%#v", req)

	return nil
}
