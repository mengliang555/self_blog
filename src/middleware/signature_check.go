package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckSignature() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Request.Header["signature"]; !ok {
			//todo panic exception
		} else {
			//todo check val is correct
		}
	}
}
