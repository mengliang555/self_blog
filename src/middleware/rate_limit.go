package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"self_blog/src/boot"
	"self_blog/src/common/common_struct"
	"sync"
)

const (
	bucketSize   = 10
	rateForToken = 10
)

var urlMap = make(map[string]*rate.Limiter)
var urlMapLock = sync.RWMutex{}

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo 限流配置支持动态调整
		if _, ok := urlMap[ctx.Request.URL.Path]; !ok {
			urlMapLock.Lock()
			urlMap[ctx.Request.URL.Path] = rate.NewLimiter(rateForToken, bucketSize)
			urlMapLock.Unlock()
		}
		if !urlMap[ctx.Request.URL.Path].Allow() {
			boot.RefLogger().Info(ctx, common_struct.RateLimitResponseStruct.Msg)
			ctx.AbortWithStatusJSON(http.StatusOK, common_struct.RateLimitResponseStruct)
			return
		}
		ctx.Next()
	}
}
