package constant_error

import "net/http"

const (
	RequestLimitMsg              = "request limit"
	ServiceUnreachable           = "service unreachable"
	RequestParamInvalidMsg       = "param invalid"
	StatusInternalServerErrorMsg = "status internal server error"
)

var DefaultError = NewError(http.StatusInternalServerError, StatusInternalServerErrorMsg)
var RequestLimitError = NewError(RequestLimitCode, RequestLimitMsg)
var ParamInvalidError = NewError(RequestParamInvalid, RequestParamInvalidMsg)
