package models

type EventType string

func (t EventType) String() string {
	return string(t)
}

type Event struct {
	Type    EventType
	Payload []byte
}
