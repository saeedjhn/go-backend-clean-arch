package mysql

type statementKey uint

const (
	StatementKeyUserCreate statementKey = iota + 1
	StatementKeyUserGetByMobile
	StatementKeyUserIsExistsByMobile
	StatementKeyUserIsExistsByEmail
	StatementKeyUserGetByID
	StatementKeyOutboxCreate
)
