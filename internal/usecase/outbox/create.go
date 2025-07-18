package outbox

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"
	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
)

func (i Interactor) Create(ctx context.Context, events []contract.DomainEvent) ([]types.ID, error) {
	ctx, span := i.trc.Span(ctx, "Create")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "Create",
		"usecase.request": events,
	})
	defer span.End()

	var idList []types.ID
	for _, event := range events {
		payload, err := event.Marshal()
		if err != nil {
			return idList, richerror.New(_opOutboxServiceCreate).
				WithErr(err).
				WithMessage(errMsgMarshal).
				WithKind(richerror.KindStatusInternalServerError)
		}

		id, err := i.repository.Create(ctx, models.OutboxEvent{
			Type:         event.GetType(),
			Payload:      payload,
			IsPublished:  false,
			ReTriedCount: 0,
		})
		if err != nil {
			return idList, err
		}

		idList = append(idList, id)
	}

	return idList, nil
}
