package entity

type Gender string

const (
	MaleGender   = Gender("male")
	FemaleGender = Gender("female")
)

var GenderStrings = map[Gender]string{
	MaleGender:   "male",
	FemaleGender: "female",
}

func (g Gender) IsValid() bool {
	_, ok := GenderStrings[g]

	return ok
}
