package user

const (
	_opUserServiceRegister     = "userservice_Register"
	_opUserServiceLogin        = "userservice_Login"
	_opUserServiceRefreshToken = "userservice_RefreshToken" // #nosec G101 // Potential hardcoded credentials

	errMsgMobileIsNotUnique            = "mobile is not unique"
	errMsgEmailIsNotUnique             = "email address is not unique"
	errMsgFailedToGeneratePasswordHash = "failed to generate password hash "
	errMsgIncorrectPassword            = "the password is incorrect"
)
