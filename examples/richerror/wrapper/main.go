package main

import (
	"github.com/rotisserie/eris"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/prettyprint"
)

type ErrorWrapper struct {
	err error
}

func NewErrorWrapper() *ErrorWrapper {
	return &ErrorWrapper{}
}

func (e *ErrorWrapper) WithErr(msg string) *ErrorWrapper {
	e.err = eris.New(msg)

	return e
}

func (e *ErrorWrapper) WithWrapErr(message string) *ErrorWrapper {
	e.err = eris.Wrap(e.err, message)
	return e
}

func (e *ErrorWrapper) ErrorTrace() map[string]interface{} {
	return eris.ToJSON(e.err, false)
}

// Repository Layer
func fetchDataFromDB() *ErrorWrapper {
	return NewErrorWrapper().WithErr("database connection failed")
}

// Service Layer
func processBusinessLogic() *ErrorWrapper {
	errWrapper := fetchDataFromDB()
	if errWrapper != nil {
		NewErrorWrapper().
			WithWrapErr("service failed to process business logic")
	}
	return errWrapper
}

// Handler Layer
func handleRequest() *ErrorWrapper {
	errWrapper := processBusinessLogic()
	if errWrapper != nil {
		NewErrorWrapper().
			WithWrapErr("handler failed to handle request")
	}
	return errWrapper
}

func main() {
	errWrapper := handleRequest()

	prettyprint.PrettyPrintData(errWrapper.ErrorTrace())
}
