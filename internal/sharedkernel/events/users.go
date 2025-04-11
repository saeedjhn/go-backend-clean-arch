package events

import (
	"encoding/json"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/google/uuid"
)

type UserCreatedEvent struct {
	EvtID            uint32   `json:"event-id"` // A Unique ID
	EvtType          string   `json:"event-type"`
	UserID           types.ID `json:"user-id"`
	EscalationReason string   `json:"escalation-reason"`
	EscalationTime   int64    `json:"escalation-time"`
	// cluster_key: Our BQ clustering key
}

func NewUserCreatedEvent(userID types.ID) *UserCreatedEvent {
	return &UserCreatedEvent{
		EvtID:            uuid.New().ID(),
		EvtType:          "users.account.created",
		UserID:           userID,
		EscalationReason: "reason",
		EscalationTime:   time.Now().Unix(),
	}
}

func (t *UserCreatedEvent) EventType() string {
	return t.EvtType
}

func (t *UserCreatedEvent) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *UserCreatedEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, t)
}
