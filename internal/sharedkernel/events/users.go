package events

import (
	"encoding/json"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/google/uuid"
)

type BasicEvent[T interface{}] struct {
	EvtID            uint32      `json:"event_id"` // A Unique ID
	EvtType          types.Event `json:"event_type"`
	EscalationReason string      `json:"escalation_reason"`
	EscalationTime   int64       `json:"escalation_time"`
	Payload          T           `json:"payload"`
}

func NewBasicEvent[T interface{}](
	eventType types.Event,
	escalationReason string,
	payload T,
) *BasicEvent[T] {
	return &BasicEvent[T]{
		EvtID:            uuid.New().ID(),
		EvtType:          eventType,
		EscalationReason: escalationReason,
		EscalationTime:   time.Now().Unix(),
		Payload:          payload,
	}
}

func (e *BasicEvent[T]) GetID() uint32 {
	return e.EvtID
}

func (e *BasicEvent[T]) GetType() types.Event {
	return e.EvtType
}

func (e *BasicEvent[T]) GetEscalationReason() string {
	return e.EscalationReason
}

func (e *BasicEvent[T]) GetEscalationTime() int64 {
	return e.EscalationTime
}

func (e *BasicEvent[T]) GetPayload() T {
	return e.Payload
}

func (e *BasicEvent[T]) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *BasicEvent[T]) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}

type UserRegisteredEvent struct {
	EvtID            uint32      `json:"event_id"` // A Unique ID
	EvtType          types.Event `json:"event_type"`
	UserID           types.ID    `json:"user_id"`
	EscalationReason string      `json:"escalation_reason"`
	EscalationTime   int64       `json:"escalation_time"`
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

func (e *UserRegisteredEvent) GetType() types.Event {
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
