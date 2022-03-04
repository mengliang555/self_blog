package constant_error

import (
	"fmt"
)

type InnerError struct {
	ErrorCode    int
	ErrorMessage string
}

func (i *InnerError) Error() string {
	return fmt.Sprintf("[code:%d][msg:%s]", i.ErrorCode, i.ErrorMessage)
}

func NewError(code int, msg string) *InnerError {
	return &InnerError{
		ErrorCode:    code,
		ErrorMessage: msg,
	}
}
