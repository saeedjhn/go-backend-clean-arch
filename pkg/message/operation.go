package message

// Rule: Op + Package(Pascal-case) + MethodOrFunction(Pascal-case).
const (

	// Validation.

	OpUserValidatorValidateRegisterRequest     = "uservalidator.ValidateRegisterRequest"
	OpUserValidatorValidateLoginRequest        = "uservalidator.ValidateLoginRequest"
	OpUserValidatorValidateRefreshTokenRequest = "uservalidator.ValidateRefreshTokenRequest"
	OpTaskValidatorValidateCreateRequest       = "taskvalidator.ValidateCreateRequest"
	OpTaskValidatorValidateCreateTaskRequest   = "taskvalidator.ValidateCreateTaskRequest"

	// UseCase.

	OpUserUsecaseRegister     = "userservice.Register"
	OpUserUsecaseLogin        = "userservice.Login"
	OpUserUsecaseRefreshToken = "userservice.RefreshToken"
	OpUserUsecaseCreateTask   = "userservice.CreateTask"
	OpTaskUsecaseCreate       = "taskservice.Create"

	// Repository.

	OpMysqlUserCreate         = "mysqluser.Create"
	OpMysqlUserIsMobileUnique = "mysqluser.IsMobileUnique"
	OpMysqlUserGetByMobile    = "mysqluser.GetByMobile"
	OpMysqlUserGetByID        = "mysqluser.GetByID"
	OpMysqlTaskCreate         = "mysqltask.Create"
	OpMysqlTaskGetByID        = "mysqltask.GetByID"
	OpMysqlTaskGetAll         = "mysqltask.GetAll"
	OpMysqlTaskGetAllByUserID = "mysqltask.GetAllByUserID"
	OpMysqlTaskIsExistsUser   = "mysqltask.IsExistsUser"

	// Compare password.

	InvalidCredentials = "invalid credentials"
)
