package entity

type Gender string

const (
	MaleGender   = Gender("male")
	FemaleGender = Gender("female")
)

var _genderStrings = map[Gender]string{ //nolint:gochecknoglobals // nothing
	MaleGender:   "male",
	FemaleGender: "female",
}

func (g Gender) IsValid() bool {
	_, ok := _genderStrings[g]

	return ok
}
