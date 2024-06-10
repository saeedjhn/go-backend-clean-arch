package message

// Rule: Op + Package(Pascal-case) + MethodOrFunction(Pascal-case)
const (
	// Validation
	OpUserValidatorValidateRegisterRequest   = "uservalidator.ValidateRegisterRequest"
	OpUserValidatorValidateLoginRequest      = "uservalidator.ValidateLoginRequest"
	OpTaskValidatorValidateCreateRequest     = "taskvalidator.ValidateCreateRequest"
	OpTaskValidatorValidateCreateTaskRequest = "taskvalidator.ValidateCreateTaskRequest"

	// UseCase
	OpUserUsecaseRegister = "userusecase.register"
	OpUserUsecaseLogin    = "userusecase.login"
	OpTaskUsecaseCreate   = "taskusecase.create"

	// Repository
	OpPqUserCreate            = "pquser.create"
	OpMysqlUserCreate         = "mysqluser.create"
	OpMysqlUserIsMobileUnique = "mysqluser.IsMobileUnique"
	OpMysqlUserGetByMobile    = "mysqluser.GetByMobile"
	OpMysqlUserGetByID        = "mysqluser.GetByID"
	OpMysqlTaskCreate         = "mysqltask.task"
	OpMysqlTaskGetByID        = "mysqltask.GetByID"
	OpMysqlTaskGetAll         = "mysqltask.GetAll"
	OpMysqlTaskGetAllByUserID = "mysqltask.GetAllByUserID"
	OpMysqlTaskIsExistsUser   = "mysqltask.IsExistsUser"

	// Compare password
	InvalidCredentials = "Invalid credentials"
)
