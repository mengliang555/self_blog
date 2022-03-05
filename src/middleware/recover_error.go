package middleware

import (
	"github.com/gin-gonic/gin"
	"self_blog/src/common/constant/constant_error"
)

func RecoverThePanicError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if val, ok := r.(*constant_error.InnerError); ok {
					_ = ctx.AbortWithError(val.ErrorCode, val)
				} else {
					_ = ctx.AbortWithError(constant_error.DefaultError.ErrorCode, constant_error.DefaultError)
				}
			}
		}()
		ctx.Next()
	}
}
