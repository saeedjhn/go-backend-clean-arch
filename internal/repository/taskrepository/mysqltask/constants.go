package mysqltask

const (
	_opMysqlTaskCreate         = "mysqltask_Create"
	_opMysqlTaskGetByID        = "mysqltask_GetByID"
	_opMysqlTaskGetAll         = "mysqltask_GetAll"
	_opMysqlTaskGetAllByUserID = "mysqltask_GetAllByUserID"
	_opMysqlTaskIsExistsUser   = "mysqltask_IsExistsUser"

	_errMsgDBRecordNotFound      = "record not found"
	_errMsgDBCantScanQueryResult = "can't scan query result"
)
