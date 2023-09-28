package routers

import (
	"GoDemo1/controller/admin"
	"GoDemo1/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	//配置中间件的两种方式
	adminRouters.Use(middleware.InitMiddleware)
	//adminRouters := r.Group("/admin", middleware.InitMiddleware)

	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/adduser", admin.UserController{}.Add)
		adminRouters.GET("/edituser", admin.EditUser)
		adminRouters.GET("/showuser", admin.UserController{}.ShowUser)
	}
}
