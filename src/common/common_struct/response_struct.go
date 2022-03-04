package common_struct

import (
	"net/http"
	"self_blog/src/common/constant/constant_error"
)

type ResponseStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var RateLimitResponseStruct = &ResponseStruct{
	Code: constant_error.RequestLimitError.ErrorCode,
	Msg:  constant_error.RequestLimitError.ErrorMessage,
	Data: nil,
}

var EmptyResponseStruct = &ResponseStruct{
	Code: http.StatusOK,
	Msg:  "Just for test,return EmptyResponseStruct",
	Data: nil,
}
