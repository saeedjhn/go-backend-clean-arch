package user

const (
	_opUserValidatorValidateRegisterRequest   = "uservalidator_ValidateRegisterRequest"
	_opUserValidatorValidateLoginRequest      = "uservalidator_ValidateLoginRequest"
	_opUserValidatorValidateProfileRequest    = "uservalidator_ValidateProfileRequest"
	_opUserValidatorValidateRefTokenRequest   = "uservalidator_validateRefTokenRequest"
	_opTaskValidatorValidateCreateTaskRequest = "uservalidator_ValidateCreateTaskRequest"

	_nameMinLen   = 3
	_nameMaxLen   = 128
	_mobileMinLen = 11
	_mobileMaxLen = 11
	_passMinLen   = 8
	_passMaxLen   = 128
	_titleMinLen  = 3
	_titleMaxLen  = 128
	_descMinLen   = 10
	_descMaxLen   = 1024

	errMsgInvalidInput = "invalid input"
)
