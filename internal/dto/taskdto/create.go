package taskdto

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateResponse struct {
}
