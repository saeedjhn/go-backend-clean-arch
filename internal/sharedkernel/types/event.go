package types

import "context"

type Event string

func (t Event) String() string {
	return string(t)
}

type EventRouterHandler func(ctx context.Context, payload []byte) error

type EventRouter map[Event]EventRouterHandler

type EventStream struct {
	Type    Event
	Payload []byte
}
