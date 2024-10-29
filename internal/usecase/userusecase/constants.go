package userusecase

const (
	_opUserServiceRegister     = "userservice_Register"
	_opUserServiceLogin        = "userservice_Login"
	_opUserServiceRefreshToken = "userservice_RefreshToken" // #nosec G101 // Potential hardcoded credentials

	_errMsgMobileIsNotUnique = "mobile is not unique"
	_errMsgIncorrectPassword = "the password is incorrect"
)
