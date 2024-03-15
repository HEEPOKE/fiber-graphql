package constants

type StatusCodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	MESSAGE_SUCCESS   = "Success"
	MESSAGE_FAIL      = "Fail"
	MESSAGE_NOTFOUND  = "Not Found"
	MESSAGE_DUPLICATE = "Duplicate"
	MESSAGE_INVALID   = "Invalid"
	CODE_SUCCESS      = "0000"
	CODE_FAILED       = "0001"
	CODE_NOT_FOUND    = "0002"
	CODE_DUPLICATE    = "0003"
	CODE_INVALID      = "0004"
)
