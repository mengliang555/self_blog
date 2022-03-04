package main

import (
	"github.com/gin-gonic/gin"
	_ "self_blog/src/middleware"

	"context"
	"self_blog/src/boot"
	_ "self_blog/src/handler/router"
	"self_blog/src/util/err_util"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	if val := boot.RefEngine().GetGinEngine().Routes(); val == nil || len(val) == 0 {
		boot.RefLogger().ErrorWithFormat(context.Background(), "no router in system %c", val)
	}
	err_util.PanicError(boot.RefEngine().GetGinEngine().Run(":8080"))
}
