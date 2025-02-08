package taskqueue

type Task struct {
	Type    string
	Payload []byte
}
