package msg

const (
	ErrorMsg500InternalServerError = "A generic error message, given when no more specific message is suitable"
	ErrorMsg501NotImplemented      = "The server either does not recognize the request method, " +
		"or it lacks the ability to fulfill the request"
	ErrorMsg502BadGateway = "The server was acting as a gateway or proxy and " +
		"received an invalid response from the upstream server"
	ErrorMsg503ServiceUnavailable = "The server is currently unavailable (overloaded or down)"
	ErrorMsg504GatewayTimeout     = "The server was acting as a gateway or " +
		"proxy and did not receive a timely response from the upstream server"
	ErrorMsg505HTTPVersionNotSupported       = "The server does not support the HTTP protocol version used in the request"
	ErrorMsg511NetworkAuthenticationRequired = ""
)
