package user

type Scanner interface {
	Scan(dest ...any) error
}
