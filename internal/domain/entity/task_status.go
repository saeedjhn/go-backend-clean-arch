package entity

/*
Since variables have a 0 default value, you should usually start your enums on a non-zero value.
There are cases where using the zero value makes sense, for example when the zero value case is the
desirable default behavior.
*/

type TaskStatus string

const (
	TaskPending    = TaskStatus("pending")
	TaskInProgress = TaskStatus("in_progress")
	TaskCompleted  = TaskStatus("completed")
)

var _taskStatusStrings = map[TaskStatus]string{ //nolint:gochecknoglobals // nothing
	TaskPending:    "pending",
	TaskInProgress: "in_progress",
	TaskCompleted:  "completed",
}

func (t TaskStatus) String() string {
	return _taskStatusStrings[t]
}

func (t TaskStatus) IsValid() bool {
	_, ok := _taskStatusStrings[t]

	return ok
}
