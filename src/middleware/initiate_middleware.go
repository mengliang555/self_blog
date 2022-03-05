package middleware

import (
	"self_blog/src/boot"
)

func init() {
	boot.RefLogger().SimpleInfo("initiate the middleware")
	boot.RefEngine().GetGinEngine().Use(RecoverThePanicError(), CheckSignature(), TraceIdSet(), RateLimit())
}
