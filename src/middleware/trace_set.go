package middleware

import (
	"github.com/gin-gonic/gin"
	"self_blog/src/util/random_util"
)

const TraceIdLength = 16

func TraceIdSet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Request.Header["trace_id"]; !ok {
			ctx.Set("trace_id", random_util.GenerateRandomHexStr(TraceIdLength))
		}
		ctx.Next()
	}
}
