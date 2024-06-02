package message

const (
	// 1XX: Information
	Msg100Continue           = "The server has received the request headers, and the client should proceed to send the request body"
	Msg101SwitchingProtocols = "The requester has asked the server to switch protocols"
	Msg102EarlyHints         = "Used with the Link header to allow the browser to start preloading resources while the server prepares a response"

	// 2XX: Successful
	Msg200Ok                          = "The request is OK"
	Msg201Created                     = "The request has been fulfilled, and a new resource is created"
	Msg202Accepted                    = "The request has been accepted for processing, but the processing has not been completed"
	Msg203NonAuthoritativeInformation = "The request has been successfully processed, but is returning information that may be from another source"
	Msg204NotContent                  = "The request has been successfully processed, but is not returning any content"
	Msg205ResetContent                = "The request has been successfully processed, but is not returning any content, and requires that the requester reset the document view"
	Msg206PartialContent              = "The server is delivering only part of the resource due to a range header sent by the client"

	// 3XX: Redirection
	Msg300MultipleChoices   = "A link list. The user can select a link and go to that location. Maximum five addresses"
	Msg301MovedPermanently  = "The requested page has moved to a new URL"
	Msg302Found             = "The requested page has moved temporarily to a new URL"
	Msg303SeeOther          = "The requested page can be found under a different URL"
	Msg304NotModified       = "Indicates the requested page has not been modified since last requested"
	Msg307TemplateRedirect  = "The requested page has moved temporarily to a new URL"
	Msg308PermanentRedirect = "The requested page has moved permanently to a new URL"

	// 4XX: Client Error
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

	// 5XX: Server Error
	ErrorMsg500InternalServerError           = "A generic error message, given when no more specific message is suitable"
	ErrorMsg501NotImplemented                = "The server either does not recognize the request method, or it lacks the ability to fulfill the request"
	ErrorMsg502BadGateway                    = "The server was acting as a gateway or proxy and received an invalid response from the upstream server"
	ErrorMsg503ServiceUnavailable            = "The server is currently unavailable (overloaded or down)"
	ErrorMsg504GatewayTimeout                = "The server was acting as a gateway or proxy and did not receive a timely response from the upstream server"
	ErrorMsg505HTTPVersionNotSupported       = "The server does not support the HTTP protocol version used in the request"
	ErrorMsg511NetworkAuthenticationRequired = ""

	ErrorMsgDBRecordNotFound      = "Record not found"
	ErrorMsgCantScanQueryResult   = "Can't scan query result"
	ErrorMsgSomethingWentWrong    = "Something went wrong"
	ErrorMsgMobileIsNotUnique     = "Mobile is not unique"
	ErrorMsgInvalidInput          = "Invalid input"
	ErrorMsgPhoneNumberIsNotValid = "Mobile is not valid"
	ErrorMsgUserNotAllowed        = "User not allowed"
	ErrorMsgCategoryIsNotValid    = "Category is not valid"

	MsgUserRegisterSuccessfully = "User is register successfully"
)
