package contract

type Topic string

func (t Topic) String() string {
	return string(t)
}

type Event struct {
	Topic   Topic
	Payload []byte
}

//go:generate mockery --name Consumer
type Consumer interface {
	Consume(chan<- Event) error
}

//go:generate mockery --name Publisher
type Publisher interface {
	Publish(event Event) error
}
