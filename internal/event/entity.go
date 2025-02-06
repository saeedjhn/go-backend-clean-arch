package event

type Topic string

func (t Topic) String() string {
	return string(t)
}

type Event struct {
	Topic   Topic
	Payload []byte
}
