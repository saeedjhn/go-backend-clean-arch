package event

//go:generate mockery --name Publisher
type Publisher interface {
	Publish(event Event) error
}

//go:generate mockery --name Consumer
type Consumer interface {
	Consume(chan<- Event) error
}
