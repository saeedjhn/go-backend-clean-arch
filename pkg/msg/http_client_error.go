package msg

const (
	ErrMsg400BadRequest   = "The request cannot be fulfilled due to bad syntax"
	ErrMsg401UnAuthorized = "The request was a legal request, but the server is refusing to respond to it. " +
		"For use when authentication is possible but has failed or not yet been provided"
	ErrMsg402PaymentRequired = "Reserved for future use"
	ErrMsg403Forbidden       = "The request was a legal request, but the server is refusing to respond to it"
	ErrMsg404NotFound        = "The requested page could not be found but may be available " +
		"again in the future"
	ErrMsg405MethodNotAllowed = "A request was made of a page using a request method not supported " +
		"by that page"
	ErrMsg406NotAcceptable               = "The server can only generate a response that is not accepted by the client"
	ErrMsg407ProxyAuthenticationRequired = "The client must first authenticate itself with the proxy"
	ErrMsg408RequestTimeout              = "The server timed out waiting for the request"
	ErrMsg409Conflict                    = "The request could not be completed because of a conflict in the request"
	ErrMsg410Gone                        = "The requested page is no longer available"
	ErrMsg411LengthRequired              = "The \"Content-Length\" is not defined. The server will not accept " +
		"the request without it"
	ErrMsg412PreconditionFailed = "The precondition given in the request evaluated to false by the server"
	ErrMsg413RequestToLarge     = "The server will not accept the request, because the request models is too large"
	ErrMsg414RequestURLToLarge  = "The server will not accept the request, because the URI is too long. " +
		"Occurs when you convert a POST request to a GET request with a long query information"
	ErrMsg415UnsupportedMediaType = "The server will not accept the request, " +
		"because the media type is not supported"
	ErrMsg416RangeNotSatisfiable = "The client has asked for a portion of the file, " +
		"but the server cannot supply that portion"
	ErrMsg417ExpectationFailed = "The server cannot meet the requirements of the Expect request-header field"
)
