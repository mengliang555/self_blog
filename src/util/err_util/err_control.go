package err_util

import (
	"context"
	"self_blog/src/boot"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorLog(ctx context.Context, err1 error, handler func(err error)) {
	boot.RefLogger().Warn(ctx, err1.Error())
	handler(err1)
}
