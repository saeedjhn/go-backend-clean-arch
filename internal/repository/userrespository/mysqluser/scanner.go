package mysqluser

type Scanner interface {
	Scan(dest ...any) error
}
