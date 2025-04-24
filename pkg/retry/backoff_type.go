package retry

type BackoffType int

const (
	ExponentialBackoff BackoffType = iota
	LinearBackoff
	ConstantBackoff
)
