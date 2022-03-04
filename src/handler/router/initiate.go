package router

import (
	"self_blog/src/boot"
	"self_blog/src/common/common_struct"
	"self_blog/src/common/constant/variable"
)

// 载入所有路由
func loadAllRouter() {
	routerMap := variable.RefRouterMap().GetInnerGroupMap
	for k, v := range routerMap() {
		group := boot.RefEngine().GetGinEngine().Group(k, v.Middleware...)
		for _, handler := range v.Handler {
			group.Handle(string(handler.RequestMethod), handler.Path, handler.Handler)
		}
	}
}

func init() {
	variable.RefRouterMap().AddGroupRouter(GroupForUserLogin, nil, []*common_struct.RequestHandler{
		loginRouter, signOutRouter, registerRouter,
	})

	variable.RefRouterMap().AddGroupRouter(GroupForUserForSelfBlog, nil, []*common_struct.RequestHandler{
		createBlogRouter, updateBlogRouter, getSelfBlogListRouter, getPartnerBlogListRouter, getBlogDetailByIdRouter,
	})

	loadAllRouter()
}
