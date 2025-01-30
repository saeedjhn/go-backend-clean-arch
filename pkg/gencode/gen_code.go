package gencode

import (
	"math/rand"
	"time"
)

func GenCode(length int, chars string) (string, error) {
	if length <= 0 {
		return "", errMustBeGTZero
	}
	if len(chars) == 0 {
		return "", errCannotEmpty
	}

	rand.NewSource(time.Now().UnixNano())
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result), nil
}

// func GenCode(length int, chars string) (string, error) {
// 	if length <= 0 {
// 		return "", errMustBeGTZero
// 	}
// 	if len(chars) == 0 {
// 		return "", errCannotEmpty
// 	}
//
// 	source := rand.NewSource(time.Now().UnixNano())
// 	rng := rand.New(source)
//
// 	result := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		result[i] = chars[rng.Intn(len(chars))]
// 	}
//
// 	return string(result), nil
// }

// GenCode generates a random string of a given length from the provided character set.
// func GenCode(length int, chars string) (string, error) {
// 	if length <= 0 {
// 		return "", errMustBeGTZero
// 	}
// 	if len(chars) == 0 {
// 		return "", errCannotEmpty
// 	}
//
// 	src := rand.NewSource(time.Now().UnixNano())
// 	rng := rand.New(src)
//
// 	var builder strings.Builder
// 	builder.Grow(length)
//
// 	for i := 0; i < length; i++ {
// 		builder.WriteByte(chars[rng.Intn(len(chars))])
// 	}
//
// 	return builder.String(), nil
// }
