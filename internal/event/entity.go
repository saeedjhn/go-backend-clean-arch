package event

type Topic string

type Event struct {
	Topic   Topic
	Payload []byte
}
