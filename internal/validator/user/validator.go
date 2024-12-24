package user

type Validator struct {
	entropyPassword float64
}

// var _ user.Validator = (*Validator)(nil)

func New(entropyPassword float64) Validator {
	return Validator{entropyPassword: entropyPassword}
}
