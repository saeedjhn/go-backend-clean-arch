package userdto

type TasksRequest struct {
	ID uint64 `param:"id" json:"id"`
}

type TasksResponse struct {
	Tasks []TaskInfo `json:"tasks"`
}
