package sanitizepkg

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"reflect"
	"regexp"
)

type Policy string

const (
	StrictPolicy    Policy = "strict_policy"
	UGCPolicy       Policy = "ugc_policy"
	StripTagsPolicy Policy = "strip_tags_policy"
	NewPolicy       Policy = "new_policy"
)

var SanitizerInstance = bluemonday.StrictPolicy()

// SanitizeStrictPolicy - bluemonday.StrictPolicy() which can be thought of as equivalent to stripping all HTML elements and their attributes as it has nothing on its allowlist.
// An example usage scenario would be blog post titles where HTML tags are not expected at all and if they are then the elements and the content of the elements
// should be stripped. This is a very strict policy.
//func sanitizeStrictPolicy(ctx fiber.Ctx) error {
//	return ctx.Next()
//}

// SanitizeUGCPolicy -  bluemonday.UGCPolicy() which allows a broad selection of HTML elements and attributes that are safe for user generated content.
// Note that this policy does not allow iframes, object, embed, styles, script, etc. An example usage scenario would be blog post bodies
// where a variety of formatting is expected along with the potential for TABLEs and IMGs.
//func sanitizeUGCPolicy(ctx fiber.Ctx) error {
//	return ctx.Next()
//}

type Sanitize struct {
	policy *bluemonday.Policy
}

func New() Sanitize {
	return Sanitize{}
}

func (s Sanitize) SetPolicy(policy Policy) Sanitize {
	switch policy {
	case StrictPolicy:
		s.policy = bluemonday.StrictPolicy()
	case UGCPolicy:
		s.policy = bluemonday.UGCPolicy()
	case StripTagsPolicy:
		s.policy = bluemonday.StripTagsPolicy()
	case NewPolicy:
		s.policy = bluemonday.NewPolicy()

	default:
		s.policy = bluemonday.StrictPolicy()
	}

	return s
}

func (s Sanitize) Sanitize(param interface{}) (interface{}, error) {
	sanitized, err := sanitizeRecursively(param)
	if err != nil {
		return nil, err
	}

	return sanitized, nil
}

func (s Sanitize) SanitizeStruct(param interface{}) (map[string]interface{}, error) {
	return sanitizeStruct(param)
}

func sanitizeRecursively(param interface{}) (interface{}, error) {
	if param == nil {
		return param, nil
	}

	paramValue := reflect.ValueOf(param)

	switch paramValue.Kind() {
	case reflect.String:
		return sanitizeString(param.(string)), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32,
		reflect.Float64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Bool:
		return param, nil

	case reflect.Slice, reflect.Array:
		return sanitizeArray(param)

	case reflect.Map:
		return sanitizeMap(param)

	case reflect.Struct:
		return sanitizeStruct(param)

	default:
		fmt.Println("type not supported", paramValue.Kind())
	}

	return nil, nil
}

func sanitizeString(param string) string {
	sanitizedHtmlStr := SanitizerInstance.Sanitize(param)

	regex := regexp.MustCompile(`\bjavascript\b`)

	return regex.ReplaceAllString(sanitizedHtmlStr, "")
}

func sanitizeArray(param interface{}) ([]interface{}, error) {

	paramValue := reflect.ValueOf(param)
	var sanitisedArray []interface{}
	for index := 0; index < paramValue.Len(); index++ {
		sanitisedParam, err := sanitizeRecursively(paramValue.Index(index).Interface())
		if err != nil {
			return nil, err
		}

		sanitisedArray = append(sanitisedArray, sanitisedParam)
	}

	return sanitisedArray, nil
}

func sanitizeMap(param interface{}) (map[string]interface{}, error) {

	paramValue := reflect.ValueOf(param)
	sanitisedMap := make(map[string]interface{})

	for _, key := range paramValue.MapKeys() {
		sanitisedParam, err := sanitizeRecursively(paramValue.MapIndex(key).Interface())
		if err != nil {
			return nil, err
		}

		sanitisedMap[key.String()] = sanitisedParam
	}

	return sanitisedMap, nil
}

func sanitizeStruct(param interface{}) (map[string]interface{}, error) {
	paramValue := reflect.ValueOf(param)
	newStruct := reflect.Indirect(paramValue)

	values := make([]interface{}, paramValue.NumField())

	sanitisedStruct := make(map[string]interface{})

	for i := 0; i < paramValue.NumField(); i++ {
		fieldName := newStruct.Type().Field(i).Name
		values[i], _ = sanitizeRecursively(paramValue.Field(i).Interface())
		sanitisedStruct[fieldName] = values[i]
	}

	return sanitisedStruct, nil
}

//func sanitizeBodyAndQuery(params interface{}) (interface{}, error) {
//	sanitisedParams, _ := sanitizeRecursively(params)
//
//	return sanitisedParams, nil
//}

//func Sanitize(out interface{}) {
//	dataMap, _ := sanitizeStruct(out)
//
//	fmt.Println("data is:", dataMap)
//
//	mapstructure.Decode(dataMap, out) // map to struct
//	//mapstructure.Decode(data, &outPtr)
//}
