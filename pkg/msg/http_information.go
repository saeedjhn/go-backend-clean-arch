package msg

const (
	Msg100Continue = "The server has received the request headers, " +
		"and the client should proceed to send the request body"
	Msg101SwitchingProtocols = "The requester has asked the server to switch protocols"
	Msg102EarlyHints         = "Used with the Link header to allow the browser to start preloading resources " +
		"while the server prepares a response"
)
