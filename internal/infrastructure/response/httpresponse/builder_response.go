package httpresponse

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse/internal"
)

type HTTPResponseBuilder struct {
	internal.HTTPResponse
}

func New() HTTPResponseBuilder {
	return HTTPResponseBuilder{internal.HTTPResponse{}}
}

func (h HTTPResponseBuilder) WithStatus(status bool) HTTPResponseBuilder {
	h.Status = status

	return h
}

func (h HTTPResponseBuilder) WithStatusCode(statusCode int) HTTPResponseBuilder {
	h.StatusCode = statusCode

	return h
}

func (h HTTPResponseBuilder) WithMessage(message string) HTTPResponseBuilder {
	h.Message = message

	return h
}

func (h HTTPResponseBuilder) WithMeta(meta interface{}) HTTPResponseBuilder {
	h.Meta = meta

	return h
}

func (h HTTPResponseBuilder) WithError(err interface{}) HTTPResponseBuilder {
	if e, ok := err.(error); ok {
		h.Meta = map[string]interface{}{"errors": e.Error()}

		return h
	}

	h.Meta = map[string]interface{}{"errors": err}

	return h
}

func (h HTTPResponseBuilder) WithData(data interface{}) HTTPResponseBuilder {
	h.Meta = map[string]interface{}{"data": data}

	return h
}

func (h HTTPResponseBuilder) Build() internal.HTTPResponse {
	return h.HTTPResponse
}

/*

func (h HTTPResponseBuilder) WithRequestID(requestID string) HTTPResponseBuilder {
	h.RequestID = requestID

	return h
}

func (h HTTPResponseBuilder) WithPath(path string) HTTPResponseBuilder {
	h.Path = path

	return h
}

func (h HTTPResponseBuilder) WithExecutionDuration(execDuration int64) HTTPResponseBuilder {
	h.ExecutionDuration = execDuration

	return h
}
*/
