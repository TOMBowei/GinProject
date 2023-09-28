package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "This is news page.")
		})
	}
}
