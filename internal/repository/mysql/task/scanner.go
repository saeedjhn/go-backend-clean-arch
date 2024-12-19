package task

type RowScanner interface {
	Scan(dest ...any) error
}

type RowsScanner interface {
	Scan(dest ...any) error
	Next() bool
}
