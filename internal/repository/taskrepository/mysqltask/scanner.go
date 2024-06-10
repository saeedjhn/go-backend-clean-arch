package mysqltask

type Scanner interface {
	Scan(dest ...any) error
}
