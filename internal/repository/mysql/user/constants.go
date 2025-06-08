package user

const (
	_opMysqlUserCreate           = "mysqluser_Create"
	_opMysqlUserIsExistsByMobile = "mysqluser_IsExistsByMobile"
	_opMysqlUserIsExistsByEmail  = "mysqluser_IsExistsByEmail"
	_opMysqlUserGetByMobile      = "mysqluser_GetByMobile"
	_opMysqlUserGetByID          = "mysqluser_GetByID"

	errMsgDBRecordNotFound      = "record not found"
	errMsgDBCantScanQueryResult = "can't scan query result"
)
