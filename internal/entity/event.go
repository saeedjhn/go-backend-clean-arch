package entity

const (
	UserRegisteredTopic = Topic("user.registered")
)

type Topic string

type RouterHandler func(event Event) error

func (t Topic) String() string {
	return string(t)
}

type Event struct {
	Topic   Topic
	Payload []byte
}
