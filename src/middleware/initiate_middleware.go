package middleware

import (
	"context"
	"self_blog/src/boot"
)

func init() {
	boot.RefLogger().Info(context.TODO(), "initiate the middleware")
	boot.RefEngine().GetGinEngine().Use(RecoverThePanicError(), CheckSignature(), TraceIdSet(), RateLimit())
}
