package message

// Rule: Op + Package(Pascal-case) + MethodOrFunction(Pascal-case)
const (
	// Validation
	OpUserValidatorValidateRegisterRequest = "uservalidator.ValidateRegisterRequest"

	// UseCase
	OpUserUsecaseRegister = "userusecase.Create"

	// Repository
	OpPqUserRegister          = "pquser.register"
	OpMysqlUserRegister       = "mysqluser.register"
	OpMysqlUserIsMobileUnique = "mysqluser.IsMobileUnique"
	OpMysqlUserGetByMobile    = "mysqluser.GetByMobile"
)
