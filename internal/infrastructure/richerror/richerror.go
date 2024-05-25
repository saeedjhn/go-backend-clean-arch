package richerror

import (
	"errors"
	. "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
)

type Op string

type EntityRichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func (e EntityRichError) GetOperation() Op {
	return e.operation
}

func (e EntityRichError) GetWrappedError() error {
	return e.wrappedError
}

func (e EntityRichError) GetMessage() string {
	return e.message
}

func (e EntityRichError) GetKind() Kind {
	return e.kind
}

func (e EntityRichError) GetMeta() map[string]interface{} {
	return e.meta
}

type RichError struct {
	EntityRichError
	//operation    Op
	//wrappedError error
	//message      string
	//kind         kind
	//meta         map[string]interface{}
}

func New(op Op) RichError {
	return RichError{EntityRichError{operation: op}}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op

	return r
}

func (r RichError) WithErr(err error) RichError {
	r.wrappedError = err

	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message

	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r RichError) Op() Op {
	if r.operation != "" {
		return r.operation
	}

	var re RichError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return ""
	}

	return re.Op()
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re RichError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) WrappedError() error {
	if r.wrappedError != nil {
		return r.wrappedError
	}

	var re RichError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return nil
	}

	return re.WrappedError()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	var re RichError
	ok := errors.As(r.wrappedError, &re)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r RichError) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	var re RichError
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return make(map[string]interface{})
	}

	return re.Meta()
}

func Analysis(err error) EntityRichError {
	var richError RichError
	switch {
	case errors.As(err, &richError):
		var re RichError
		errors.As(err, &re)

		return EntityRichError{
			operation:    re.Op(),
			wrappedError: re.WrappedError(),
			message:      re.Message(),
			kind:         re.Kind(),
			meta:         re.Meta(),
		}

	default:
		return EntityRichError{}
	}
}
