package message

// Rule: Op + Package(Pascal-case) + MethodOrFunction(Pascal-case)
const (
	// Validation
	OpUserValidatorValidateRegisterRequest = "uservalidator.ValidateRegisterRequest"
	OpUserValidatorValidateLoginRequest    = "uservalidator.ValidateLoginRequest"

	// UseCase
	OpUserUsecaseRegister = "userusecase.register"
	OpUserUsecaseLogin    = "userusecase.login"

	// Repository
	OpPqUserRegister          = "pquser.register"
	OpMysqlUserRegister       = "mysqluser.register"
	OpMysqlUserIsMobileUnique = "mysqluser.IsMobileUnique"
	OpMysqlUserGetByMobile    = "mysqluser.GetByMobile"
	OpMysqlUserGetByID        = "mysqluser.GetByID"

	// Compare password
	InvalidCredentials = "Invalid credentials"
)
