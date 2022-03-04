package constant_error

import "net/http"

const (
	RequestLimitMsg    = "request limit"
	ServiceUnreachable = "service unreachable"
)

var DefaultError = &InnerError{http.StatusInternalServerError, "StatusInternalServerError"}
var RequestLimitError = &InnerError{RequestLimitCode, RequestLimitMsg}
