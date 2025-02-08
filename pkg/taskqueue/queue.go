package taskqueue

const (
	QueueCritical Queue = "critical"
	QueueDefault  Queue = "default"
	QueueLow      Queue = "low"

	_criticalValue = 6
	_defaultValue  = 3
	_lowValue      = 1
)

var _queues = map[string]int{ //nolint:gochecknoglobals // nothing
	string(QueueCritical): _criticalValue,
	string(QueueDefault):  _defaultValue,
	string(QueueLow):      _lowValue,
}

type Queue string

func (q Queue) String() string {
	return string(q)
}
