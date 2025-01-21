package msg

const (
	Msg200Ok       = "The request is OK"
	Msg201Created  = "The request has been fulfilled, and a new resource is created"
	Msg202Accepted = "The request has been accepted for processing, " +
		"but the processing has not been completed"
	Msg203NonAuthoritativeInformation = "The request has been successfully processed, " +
		"but is returning information that may be from another source"
	Msg204NotContent   = "The request has been successfully processed, but is not returning any content"
	Msg205ResetContent = "The request has been successfully processed, " +
		"but is not returning any content, and requires that the requester reset the document view"
	Msg206PartialContent = "The server is delivering only part of the resource due to a range " +
		"header sent by the client"
)
