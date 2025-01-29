package gencode

import "errors"

var (
	errCannotEmpty  = errors.New("chars cannot be empty")
	errMustBeGTZero = errors.New("length must be greater than zero")
)
