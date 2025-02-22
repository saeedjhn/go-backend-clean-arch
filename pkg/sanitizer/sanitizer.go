package sanitizer

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Sanitizer struct {
	strategies []SanitizationStrategy
}

func New() *Sanitizer {
	return &Sanitizer{}
}

func (s *Sanitizer) AddStrategy(strategy SanitizationStrategy) {
	s.strategies = append(s.strategies, strategy)
}

func (s *Sanitizer) ApplyStrategies(input string) (string, error) {
	result := input
	for _, strategy := range s.strategies {
		var err error
		// result = strategy.Sanitize(result)
		result, err = strategy.Sanitize(result)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

func (s *Sanitizer) String(input string) (string, error) {
	return s.ApplyStrategies(input)
}

func (s *Sanitizer) Array(param interface{}) ([]interface{}, error) {
	paramValue := reflect.ValueOf(param)
	var sanitisedArray []interface{}

	//nolint:intrange // for loop can be changed to use an integer range (Go 1.22+)
	for index := 0; index < paramValue.Len(); index++ {
		sanitisedParam, err := s.recursively(paramValue.Index(index).Interface())
		if err != nil {
			return nil, err
		}

		sanitisedArray = append(sanitisedArray, sanitisedParam)
	}

	return sanitisedArray, nil
}

func (s *Sanitizer) Map(param interface{}) (map[string]interface{}, error) {
	paramValue := reflect.ValueOf(param)
	sanitisedMap := make(map[string]interface{})

	for _, key := range paramValue.MapKeys() {
		sanitisedParam, err := s.recursively(paramValue.MapIndex(key).Interface())
		if err != nil {
			return nil, err
		}

		sanitisedMap[key.String()] = sanitisedParam
	}

	return sanitisedMap, nil
}

func (s *Sanitizer) StructToMap(param interface{}) (map[string]interface{}, error) {
	paramValue := reflect.ValueOf(param)
	newStruct := reflect.Indirect(paramValue)

	values := make([]interface{}, paramValue.NumField())

	sanitisedStruct := make(map[string]interface{})

	//nolint:intrange // for loop can be changed to use an integer range (Go 1.22+)
	for i := 0; i < paramValue.NumField(); i++ {
		fieldName := newStruct.Type().Field(i).Name
		values[i], _ = s.recursively(paramValue.Field(i).Interface())
		sanitisedStruct[fieldName] = values[i]
	}

	return sanitisedStruct, nil
}

func (s *Sanitizer) Struct(ptr interface{}) error {
	if !IsPointer(ptr) {
		return errNonPointer
	}

	dataMap, err := s.StructToMap(reflect.ValueOf(ptr).Elem().Interface())
	if err != nil {
		return fmt.Errorf("failed to convert struct to map: %w", err)
	}

	// Go library for decoding generic map values into native Go structures and vice versa.
	err = mapstructure.Decode(dataMap, ptr)
	if err != nil {
		return fmt.Errorf("failed to decode map into struct: %w", err)
	}

	return nil
}

func (s *Sanitizer) Any(param interface{}) (interface{}, error) {
	sanitized, err := s.recursively(param)
	if err != nil {
		return nil, err
	}

	return sanitized, nil
}

func (s *Sanitizer) recursively(param interface{}) (interface{}, error) {
	if param == nil {
		return param, nil
	}

	paramValue := reflect.ValueOf(param)
	switch paramValue.Kind() { //nolint:exhaustive // missing cases in switch of type paramValue.Kind()
	case reflect.String:
		return s.String(reflect.ValueOf(param).String())
		// return s.String(reflect.ValueOf(param).String()), nil
		// return s.String(param.(string)), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32,
		reflect.Float64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool:
		return param, nil

	case reflect.Slice, reflect.Array:
		return s.Array(param)

	case reflect.Map:
		return s.Map(param)

	case reflect.Struct:
		return s.StructToMap(param)

	default:
		return nil, fmt.Errorf("type not supported %v", paramValue.Kind().String())
	}
}
