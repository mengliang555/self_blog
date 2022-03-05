package common_struct

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"self_blog/src/boot"
	"self_blog/src/common/constant/base_constant"
	"self_blog/src/common/constant/constant_error"
)

type RequestHandler struct {
	RequestMethod base_constant.HttpMethod
	Group         string
	Path          string
	Param         interface{}
	Handler       func(ctx *gin.Context)
}

func (r *RequestHandler) ParamCheck(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ctx.BindJSON(r.Param)
		ans, err := ioutil.ReadAll(ctx.Request.Body)
		boot.RefLogger().SimpleInfo(string(ans))
		if err != nil {
			boot.RefLogger().SimpleInfo(err.Error())
			panic(constant_error.ParamInvalidError)
		}
		handler(ctx)
	}
}

func (r *RequestHandler) Build() *RequestHandler {
	r.Handler = r.ParamCheck(r.Handler)
	return r
}
