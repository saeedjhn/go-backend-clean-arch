package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/events"
)

func (h Handler) Registered(ctx context.Context, payload []byte) error {
	var ur events.UserRegisteredEvent
	if err := ur.Unmarshal(payload); err != nil {
		return err
	}

	if err := h.userIntr.RegisteredHandler(ctx, ur); err != nil {
		return err
	}

	return nil
}
