package common_struct

import (
	"github.com/gin-gonic/gin"
	"self_blog/src/common/constant/base_constant"
)

type RequestHandler struct {
	RequestMethod base_constant.HttpMethod
	Group         string
	Path          string
	Handler       func(ctx *gin.Context)
}
