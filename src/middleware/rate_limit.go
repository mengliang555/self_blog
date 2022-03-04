package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"self_blog/src/common/common_struct"
	"sync"
)

const (
	bucketSize   = 10
	rateForToken = 1
)

var urlMap = make(map[string]*rate.Limiter)
var urlMapLock = sync.RWMutex{}

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo 限流配置支持动态调整
		if val, ok := urlMap[ctx.Request.URL.Path]; ok {
			if !val.Allow() {
				ctx.AbortWithStatusJSON(http.StatusOK, common_struct.RateLimitResponseStruct)
				return
			}
		} else {
			urlMapLock.Lock()
			urlMap[ctx.Request.URL.Path] = rate.NewLimiter(rateForToken, bucketSize)
			urlMapLock.Unlock()
		}
		ctx.Next()
	}
}
