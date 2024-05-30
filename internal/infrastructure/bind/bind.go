package bind

import (
	"encoding/json"
	"errors"
)

func CheckErrFromBind(err error) error {
	var v *json.UnmarshalTypeError
	if errors.As(err, &v) {
		return NewJsonUnmarshalTypeErr(v)
	}

	return err
}
