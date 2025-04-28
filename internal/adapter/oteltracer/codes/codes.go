package codes

import (
	"fmt"
)

const (
	// Unset is the default status code.
	Unset Code = 0

	// Error indicates the operation contains an error.
	//
	// NOTE: The error code in OTLP is 2.
	// The value of this enum is only relevant to the internals
	// of the Go SDK.
	Error Code = 1

	// Ok indicates operation has been validated by an Application developers
	// or Operator to have completed successfully, or contain no error.
	//
	// NOTE: The Ok code in OTLP is 1.
	// The value of this enum is only relevant to the internals
	// of the Go SDK.
	Ok Code = 2
)

// Code is an 32-bit representation of a status state.
type Code uint32

// Uint32 returns the Code as a uint32.
func (c Code) Uint32() uint32 {
	return uint32(c)
}

// String returns the Code as a string.
func (c Code) String() string {
	return fmt.Sprintf("%d", uint32(c)) //nolint:perfsprint // nothing
}
