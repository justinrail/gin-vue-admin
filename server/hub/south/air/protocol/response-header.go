package protocol

type ResponseHeader struct {
	requestType string
	errorCode   string
	errmsg      string
}
