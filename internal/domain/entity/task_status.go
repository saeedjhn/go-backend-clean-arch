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

var _taskStatusStrings = map[TaskStatus]string{
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

// type TaskStatus uint

// const (
// 	TaskPending TaskStatus = iota + 1
// 	TaskInProgress
// 	TaskCompleted
// )

// const (
// 	TaskPendingStr    = "pending"
// 	TaskInProgressStr = "in_progress"
// 	TaskCompletedStr  = "completed"
// )

// func (s TaskStatus) String() string {
// 	switch s {
// 	case TaskPending:
// 		return TaskPendingStr
// 	case TaskInProgress:
// 		return TaskInProgressStr
// 	case TaskCompleted:
// 		return TaskCompletedStr
// 	}
//
// 	return ""
// }

// func MapToTaskStatus(status string) TaskStatus {
// 	switch status {
// 	case TaskPendingStr:
// 		return TaskPending
// 	case TaskInProgressStr:
// 		return TaskInProgress
// 	case TaskCompletedStr:
// 		return TaskCompleted
// 	}
//
// 	return TaskStatus(0)
// }
