package variable

import (
	"github.com/gin-gonic/gin"
	"self_blog/src/common/common_struct"
)

type routerMap struct {
	innerGroupMap map[string]*GroupRouterMap
}

type GroupRouterMap struct {
	Handler    []*common_struct.RequestHandler
	Middleware []gin.HandlerFunc
}

var routerMapData *routerMap

func init() {
	routerMapData = &routerMap{
		innerGroupMap: make(map[string]*GroupRouterMap),
	}
}

func RefRouterMap() *routerMap {
	return routerMapData
}

func (r *routerMap) GetInnerGroupMap() map[string]*GroupRouterMap {
	return r.innerGroupMap
}

func (r *routerMap) AddGroupRouter(group string, middleware []gin.HandlerFunc, handler []*common_struct.RequestHandler) *routerMap {
	if handler == nil || len(handler) == 0 {
		return r
	}
	if _, ok := r.innerGroupMap[group]; !ok {
		r.innerGroupMap[group] = &GroupRouterMap{
			Handler:    make([]*common_struct.RequestHandler, 0),
			Middleware: make([]gin.HandlerFunc, 0),
		}
	}
	r.innerGroupMap[group].Middleware = append(r.innerGroupMap[group].Middleware, middleware...)
	r.innerGroupMap[group].Handler = append(r.innerGroupMap[group].Handler, handler...)
	return r
}
