package taskdto

type FindAllByUserIDRequest struct {
	UserID uint64 `param:"id" json:"user_id"`
}

type FindAllByUserIDResponse struct {
	Data []TaskInfo `json:"data"`
}
