package common_struct

import (
	"github.com/gin-gonic/gin"
	"self_blog/src/common/constant/variable"
)

type RequestHandler struct {
	RequestMethod variable.HttpMethod
	Group         string
	Path          string
	Handler       func(ctx *gin.Context)
}
