package msg

const (
	ErrMsgSomethingWentWrong   = "something went wrong"
	ErrMsgCantPrepareStatement = "failed to prepare SQL statement"
	IncorrectPassword          = "the password is incorrect"

	MsgLoggedIn    = "You have logged in successfully"
	MsgRegister    = "Registration completed successfully"
	MsgProfileSeen = "Your profile information has been retrieved"

	MsgCreated = "New resource created successfully"
	MsgRead    = "Requested resource loaded successfully"
	MsgUpdate  = "The resource has been updated successfully"
	MsgDelete  = "The resource has been deleted successfully"

	MsgRefreshTokenRecreated = "Your session has been extended"
)
