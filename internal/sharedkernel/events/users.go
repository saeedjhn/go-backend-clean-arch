package events

import (
	"encoding/json"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/google/uuid"
)

type UserRegisteredEvent struct {
	EvtID            uint32           `json:"event_id"` // A Unique ID
	EvtType          models.EventType `json:"event_type"`
	UserID           types.ID         `json:"user_id"`
	EscalationReason string           `json:"escalation_reason"`
	EscalationTime   int64            `json:"escalation_time"`
	// cluster_key: Our BQ clustering key
}

func NewUserRegisteredEvent(userID types.ID, reason string) *UserRegisteredEvent {
	return &UserRegisteredEvent{
		EvtID:            uuid.New().ID(),
		EvtType:          UsersRegistered,
		UserID:           userID,
		EscalationReason: reason,
		EscalationTime:   time.Now().Unix(),
	}
}

func (e *UserRegisteredEvent) GetID() uint32 {
	return e.EvtID
}

func (e *UserRegisteredEvent) GetType() models.EventType {
	return e.EvtType
}

func (e *UserRegisteredEvent) GetEscalationReason() string {
	return e.EscalationReason
}

func (e *UserRegisteredEvent) GetEscalationTime() int64 {
	return e.EscalationTime
}

func (e *UserRegisteredEvent) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserRegisteredEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
