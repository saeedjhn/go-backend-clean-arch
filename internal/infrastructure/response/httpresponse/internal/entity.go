package internal

type HTTPResponse struct {
	Status            bool          `json:"Status"`
	StatusCode        int           `json:"status_code"`
	RequestID         string        `json:"request_id"`
	Path              string        `json:"Path"`
	ExecutionDuration string        `json:"execution_duration"`
	Message           []string      `json:"Message"`
	Meta              []interface{} `json:"Meta"`
}
