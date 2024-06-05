package bind

import (
	"encoding/json"
	"fmt"
)

type JSONUnmarshalTypeErr struct {
	jsonTypeErr *json.UnmarshalTypeError
}

func NewJSONUnmarshalTypeErr(jsonTypeErr *json.UnmarshalTypeError) *JSONUnmarshalTypeErr {
	return &JSONUnmarshalTypeErr{jsonTypeErr: jsonTypeErr}
}

func (c JSONUnmarshalTypeErr) Error() string {
	// v.Field: fmt.Sprintf("cannot convert %s for name of type %s", v.Value, v.Type),
	return fmt.Sprintf(
		"%s: cannot %s for name of the %s",
		c.jsonTypeErr.Field,
		c.jsonTypeErr.Value,
		c.jsonTypeErr.Type,
	)
}
