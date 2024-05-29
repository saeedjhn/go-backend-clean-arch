package httpresponse

type HTTPResponse struct {
	Status            bool        `json:"status"`
	StatusCode        int         `json:"status_code"`
	RequestID         string      `json:"request_id"`
	Path              string      `json:"Path"`
	ExecutionDuration int64       `json:"execution_duration"`
	Message           string      `json:"message"`
	Meta              interface{} `json:"meta"`
}

/*
more:
*/
