package user

const (
	_opUserServiceRegister     = "userservice_Register"
	_opUserServiceLogin        = "userservice_Login"
	_opUserServiceRefreshToken = "userservice_RefreshToken" // #nosec G101 // Potential hardcoded credentials

	errMsgMobileIsNotUnique = "mobile is not unique"
	errMsgIncorrectPassword = "the password is incorrect"
)
