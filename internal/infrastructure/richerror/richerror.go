package richerror

import (
	"errors"
	"github.com/rotisserie/eris"
	. "go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
)

type Op string

type RichErr struct {
	op           Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
	stackTrace   map[string]interface{}
}

func (e RichErr) Op() Op {
	return e.op
}

func (e RichErr) WrappedError() error {
	return e.wrappedError
}

func (e RichErr) Error() string {
	var err error
	if errors.As(e.wrappedError, &err) {
		return e.wrappedError.Error()
	}

	return ""
}

func (e RichErr) Message() string {
	return e.message
}

func (e RichErr) Kind() Kind {
	return e.kind
}

func (e RichErr) Meta() map[string]interface{} {
	return e.meta
}

func (e RichErr) StackTrace() map[string]interface{} {
	return e.stackTrace
}

func (e RichErr) Get() map[string]interface{} {
	return map[string]interface{}{
		"op":          e.Op(),
		"error":       e.Error(),
		"message":     e.Message(),
		"kind":        e.Kind(),
		"meta":        e.Meta(),
		"stack_trace": e.StackTrace(),
	}
}

type RichErrBuilder struct {
	RichErr
}

func New(op Op) RichErrBuilder {
	return RichErrBuilder{RichErr{op: op}}
}

func (r RichErrBuilder) WithOp(op Op) RichErrBuilder {
	r.op = op

	return r
}

func (r RichErrBuilder) WithErr(err error) RichErrBuilder {
	r.wrappedError = err

	return r
}

func (r RichErrBuilder) WithMessage(message string) RichErrBuilder {
	r.message = message

	return r
}

func (r RichErrBuilder) WithKind(kind Kind) RichErrBuilder {
	r.kind = kind

	return r
}

func (r RichErrBuilder) WithMeta(meta map[string]interface{}) RichErrBuilder {
	r.meta = meta

	return r
}

func (r RichErrBuilder) WithStackTrace(message string) RichErrBuilder {
	e := eris.ToJSON(eris.New(message), true)
	r.stackTrace = e

	return r
}

func (r RichErrBuilder) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r RichErrBuilder) Op() Op {
	if r.op != "" {
		return r.op
	}

	var re RichErrBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return ""
	}

	return re.Op()
}

func (r RichErrBuilder) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re RichErrBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichErrBuilder) WrappedError() error {
	if r.wrappedError != nil {
		return r.wrappedError
	}

	var re RichErrBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return nil
	}

	return re.WrappedError()
}

func (r RichErrBuilder) Message() string {
	if r.message != "" {
		return r.message
	}

	var re RichErrBuilder
	ok := errors.As(r.wrappedError, &re)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r RichErrBuilder) Meta() map[string]interface{} {
	if len(r.meta) != 0 {
		return r.meta
	}

	var re RichErrBuilder
	ok := errors.As(r.wrappedError, &re)
	if !ok {
		return make(map[string]interface{})
	}

	return re.Meta()
}

func (r RichErrBuilder) Build() RichErr {
	return r.RichErr
}

func Analysis(err error) (RichErr, error) {
	var richError RichErrBuilder
	switch {
	case errors.As(err, &richError):
		var re RichErrBuilder
		errors.As(err, &re)

		return RichErr{
			op:           re.Op(),
			wrappedError: re.WrappedError(),
			message:      re.Message(),
			kind:         re.Kind(),
			meta:         re.Meta(),
			stackTrace:   re.StackTrace(),
		}, nil

	default:
		return RichErr{}, err
	}
}
