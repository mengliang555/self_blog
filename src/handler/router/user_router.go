package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self_blog/src/common/common_struct"
	"self_blog/src/common/constant/base_constant"
)

// user behave
//------------------------------------------------------------------------------------------------------------------
const (
	GroupForUserLogin = "group_for_user_for_login"
	login             = "login"
	signOut           = "sign_out"
	register          = "register"
)

var loginRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.POST,
	Group:         GroupForUserLogin,
	Path:          login,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

var signOutRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.POST,
	Group:         GroupForUserLogin,
	Path:          signOut,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

var registerRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.POST,
	Group:         GroupForUserLogin,
	Path:          register,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

//------------------------------------------------------------------------------------------------------------------
const (
	GroupForUserForSelfBlog = "group_for_user_for_self_blog"
	createBlog              = "create_blog"
	updateBlog              = "update_blog"
	getSelfBlogList         = "get_self_blog_list"
	getPartnerBlogList      = "get_partner_blog_list"
	getBlogDetailById       = "get_blog_detail_by_id"
)

var createBlogRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.POST,
	Group:         GroupForUserForSelfBlog,
	Path:          createBlog,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

var updateBlogRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.PUT,
	Group:         GroupForUserForSelfBlog,
	Path:          updateBlog,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

var getSelfBlogListRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.GET,
	Group:         GroupForUserForSelfBlog,
	Path:          getSelfBlogList,
	// todo 使用handler 下的method包相关方法
	Handler: func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common_struct.EmptyResponseStruct)
	},
}

var getPartnerBlogListRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.GET,
	Group:         GroupForUserForSelfBlog,
	Path:          getPartnerBlogList,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}

var getBlogDetailByIdRouter = &common_struct.RequestHandler{
	RequestMethod: base_constant.GET,
	Group:         GroupForUserForSelfBlog,
	Path:          getBlogDetailById,
	// todo 使用handler 下的method包相关方法
	Handler: nil,
}
