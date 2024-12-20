package task

type Validator struct {
}

// var _ task.Validator = (*Validator)(nil)

func New() Validator {
	return Validator{}
}
