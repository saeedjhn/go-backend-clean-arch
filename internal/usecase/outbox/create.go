package outbox

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

func (i *Interactor) Create(ctx context.Context, events []contract.DomainEvent) error {
	ctx, span := i.trc.Span(ctx, "Create")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Create",
		"usecase.request": events,
	})
	defer span.End()

	for _, event := range events {
		payload, err := event.Marshal()
		if err != nil {
			return richerror.New(_opOutboxServiceCreate).
				WithErr(err).
				WithMessage(errMsgMarshal).
				WithKind(richerror.KindStatusInternalServerError)
		}

		if err = i.repository.Create(ctx, models.OutboxEvent{
			Type:         event.GetType(),
			Payload:      payload,
			IsPublished:  false,
			ReTriedCount: 0,
		}); err != nil {
			return err
		}
	}

	return nil
}
