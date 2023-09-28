package routers

import (
	"GoDemo1/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/adduser", admin.UserController{}.Add)
		adminRouters.GET("/edituser", admin.EditUser)
	}
}
