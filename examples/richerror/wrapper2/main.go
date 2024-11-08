package main

import (
	"errors"
	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/prettyprint"
)

type AppError struct {
	Code    string
	Message string
	Err     error
	WrapErr error
}

func NewAppError() *AppError {
	return &AppError{}
	//var err *AppError
	//iErr := innerErr
	//
	//if errors.As(innerErr, &err) {
	//	iErr = err.WrapErr
	//}
	//
	//return &AppError{
	//	Code:    code,
	//	Message: message,
	//	WrapErr: eris.Wrap(iErr, message),
	//}
}

func (e *AppError) WithErr(err error) *AppError {
	e.Err = err
	e.WrapErr = eris.New(err.Error())

	return e
}

func (e *AppError) WithWrapErr(err error, msg string) *AppError {
	var appErr *AppError
	ei := err

	if errors.As(err, &appErr) {
		ei = appErr.WrapErr
	}

	e.WrapErr = eris.Wrap(ei, msg)

	return e
}

func (e *AppError) Error() string {
	return e.Err.Error()
	//return fmt.Sprintf("[[%s]] %s: %v", e.Code, e.Message, e.WrapErr)
}

func (e *AppError) StackTrace() map[string]interface{} {
	return eris.ToJSON(e.WrapErr, false)
}

func (e *AppError) Unwrap() error {
	return e.WrapErr
}

func main() {
	notFoundErr := errors.New("external")
	e := NewAppError().WithErr(notFoundErr)
	ee := NewAppError().WithWrapErr(e, "msg1")
	eee := NewAppError().WithWrapErr(ee, "msg2")

	prettyprint.PrettyPrintData(eee.StackTrace())
}
