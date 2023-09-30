package routers

import (
	"GoDemo1/controller/index"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
		defaultRouters.GET("/news", index.IndexController{}.News)
		defaultRouters.GET("/news2", index.IndexController{}.News2)
		defaultRouters.GET("/cookie", index.IndexController{}.GetCookie)
		defaultRouters.GET("/deletecookie", index.IndexController{}.DeleteCookie)

	}
}
