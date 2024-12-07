package token

type Token struct {
	config Config
}

func New(config Config) *Token {
	return &Token{config: config}
}
