package internal

type HTTPResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Meta       interface{} `json:"meta"`
}

/*
more:
	//RequestID         string      `json:"request_id"`
	//Path              string      `json:"Path"`
	//ExecutionDuration int64       `json:"execution_duration"`
*/
