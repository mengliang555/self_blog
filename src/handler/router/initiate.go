package router

import (
	"self_blog/src/boot"
	"self_blog/src/common/constant/variable"
)

func init() {
	//todo 自定义中间件执行
	routerMap := variable.RefRouterMap().GetInnerGroupMap
	for k, v := range routerMap() {
		group := boot.RefEngine().GetGinEngine().Group(k, v.Middleware...)
		for _, handler := range v.Handler {
			group.Handle(string(handler.RequestMethod), handler.Path, handler.Handler)
		}
	}
}
