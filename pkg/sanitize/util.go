package sanitize

import (
	"reflect"
)

// var SanitizerInstance = bluemonday.StrictPolicy()
//
//	func sanitizeRecursively(param interface{}) (interface{}, error) {
//		if param == nil {
//			return param, nil
//		}
//
//		paramValue := reflect.ValueOf(param)
//
//		switch paramValue.Kind() {
//		case reflect.String:
//			return sanitizeString(param.(string)), nil
//
//		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32,
//			reflect.Float64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool:
//			return param, nil
//
//		case reflect.Slice, reflect.Array:
//			return sanitizeArray(param)
//
//		case reflect.Map:
//			return sanitizeMap(param)
//
//		case reflect.Struct:
//			return sanitizeStruct(param)
//
//		default:
//			log.Println("type not supported", paramValue.Kind())
//		}
//
//		return nil, nil
//	}
//
//	func sanitizeString(param string) string {
//		sanitizedHtmlStr := SanitizerInstance.Sanitize(param)
//
//		regex := regexp.MustCompile(`\bjavascript\b`)
//
//		return regex.ReplaceAllString(sanitizedHtmlStr, "")
//	}
//
// func sanitizeArray(param interface{}) ([]interface{}, error) {
//
//		paramValue := reflect.ValueOf(param)
//		var sanitisedArray []interface{}
//		for index := 0; index < paramValue.Len(); index++ {
//			sanitisedParam, err := sanitizeRecursively(paramValue.Index(index).Interface())
//			if err != nil {
//				return nil, err
//			}
//
//			sanitisedArray = append(sanitisedArray, sanitisedParam)
//		}
//
//		return sanitisedArray, nil
//	}
//
// func sanitizeMap(param interface{}) (map[string]interface{}, error) {
//
//		paramValue := reflect.ValueOf(param)
//		sanitisedMap := make(map[string]interface{})
//
//		for _, key := range paramValue.MapKeys() {
//			sanitisedParam, err := sanitizeRecursively(paramValue.MapIndex(key).Interface())
//			if err != nil {
//				return nil, err
//			}
//
//			sanitisedMap[key.String()] = sanitisedParam
//		}
//
//		return sanitisedMap, nil
//	}
//
//	func sanitizeStruct(param interface{}) (map[string]interface{}, error) {
//		paramValue := reflect.ValueOf(param)
//		newStruct := reflect.Indirect(paramValue)
//
//		values := make([]interface{}, paramValue.NumField())
//
//		sanitisedStruct := make(map[string]interface{})
//
//		for i := 0; i < paramValue.NumField(); i++ {
//			fieldName := newStruct.Type().Field(i).Name
//			values[i], _ = sanitizeRecursively(paramValue.Field(i).Interface())
//			sanitisedStruct[fieldName] = values[i]
//		}
//
//		return sanitisedStruct, nil
//	}
//
// func sanitizeBodyAndQuery(params interface{}) (interface{}, error) {
//	sanitisedParams, _ := sanitizeRecursively(params)
//
//	return sanitisedParams, nil
// }

// func mapToStruct(ptr interface{}) error {
//	if !isPointer(ptr) {
//		return fmt.Errorf("please give me the pointer arg")
//	}
//
//	dataMap, err := sanitizeStruct(reflect.ValueOf(ptr).Elem().Interface())
//	//dataMap, err := sanitizeStruct(ptr)
//	if err != nil {
//		return fmt.Errorf("cannot perform the operation")
//	}
//	// Go library for decoding generic map values into native Go structures and vice versa.
//	err = mapstructure.Decode(dataMap, &ptr)
//	if err != nil {
//		return fmt.Errorf("in decode, %w", err)
//	}
//
//	return nil
// }

func isPointer(param interface{}) bool {
	return reflect.ValueOf(param).Kind() == reflect.Ptr
}
