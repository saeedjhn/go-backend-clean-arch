package message

const (
	ErrorMsg400BadRequest                  = "The request cannot be fulfilled due to bad syntax"
	ErrorMsg401UnAuthorized                = "The request was a legal request, but the server is refusing to respond to it. For use when authentication is possible but has failed or not yet been provided"
	ErrorMsg402PaymentRequired             = "Reserved for future use"
	ErrorMsg403Forbidden                   = "The request was a legal request, but the server is refusing to respond to it"
	ErrorMsg404NotFound                    = "The requested page could not be found but may be available again in the future"
	ErrorMsg405MethodNotAllowed            = "A request was made of a page using a request method not supported by that page"
	ErrorMsg406NotAcceptable               = "The server can only generate a response that is not accepted by the client"
	ErrorMsg407ProxyAuthenticationRequired = "The client must first authenticate itself with the proxy"
	ErrorMsg408RequestTimeout              = "The server timed out waiting for the request"
	ErrorMsg409Conflict                    = "The request could not be completed because of a conflict in the request"
	ErrorMsg410Gone                        = "The requested page is no longer available"
	ErrorMsg411LengthRequired              = `The "Content-Length" is not defined. The server will not accept the request without it`
	ErrorMsg412PreconditionFailed          = "The precondition given in the request evaluated to false by the server"
	ErrorMsg413RequestToLarge              = "The server will not accept the request, because the request entity is too large"
	ErrorMsg414RequestURLToLarge           = "The server will not accept the request, because the URI is too long. Occurs when you convert a POST request to a GET request with a long query information"
	ErrorMsg415UnsupportedMediaType        = "The server will not accept the request, because the media type is not supported"
	ErrorMsg416RangeNotSatisfiable         = "The client has asked for a portion of the file, but the server cannot supply that portion"
	ErrorMsg417ExpectationFailed           = "The server cannot meet the requirements of the Expect request-header field"
)
