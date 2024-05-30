package bind

import (
	"encoding/json"
	"fmt"
)

type JsonUnmarshalTypeErr struct {
	jsonTypeErr *json.UnmarshalTypeError
}

func NewJsonUnmarshalTypeErr(jsonTypeErr *json.UnmarshalTypeError) *JsonUnmarshalTypeErr {
	return &JsonUnmarshalTypeErr{jsonTypeErr: jsonTypeErr}
}

func (c JsonUnmarshalTypeErr) Error() string {
	// v.Field: fmt.Sprintf("cannot convert %s for name of type %s", v.Value, v.Type),
	return fmt.Sprintf(
		"%s: cannot %s for name of the %s",
		c.jsonTypeErr.Field,
		c.jsonTypeErr.Value,
		c.jsonTypeErr.Type,
	)
}
