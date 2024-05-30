package httpresponse

import "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse/internal"

type Builder struct {
	internal.HTTPResponse
}

func New() Builder {
	return Builder{internal.HTTPResponse{}}
}

func (h Builder) WithStatus(status bool) Builder {
	h.Status = status

	return h
}

func (h Builder) WithStatusCode(statusCode int) Builder {
	h.StatusCode = statusCode

	return h
}

func (h Builder) WithMessage(message string) Builder {
	h.Message = message

	return h
}

func (h Builder) WithMeta(meta interface{}) Builder {
	h.Meta = meta

	return h
}

func (h Builder) WithError(err interface{}) Builder {
	h.Meta = map[string]interface{}{"errors": err}

	return h
}

func (h Builder) WithData(data interface{}) Builder {
	h.Meta = map[string]interface{}{"data": data}

	return h
}

func (h Builder) Build() internal.HTTPResponse {
	return h.HTTPResponse
}

/*

func (h Builder) WithRequestID(requestID string) Builder {
	h.RequestID = requestID

	return h
}

func (h Builder) WithPath(path string) Builder {
	h.Path = path

	return h
}

func (h Builder) WithExecutionDuration(execDuration int64) Builder {
	h.ExecutionDuration = execDuration

	return h
}
*/
