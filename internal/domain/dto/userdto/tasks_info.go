package userdto

type TasksRequest struct {
	ID uint `param:"id" json:"id"`
}

type TasksResponse struct {
	Tasks []TaskInfo `json:"tasks"`
}
