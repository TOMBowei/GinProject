package routers

import (
	"GoDemo1/controller/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/addArticle", api.ApiController{}.AddArticle)
		apiRouters.GET("/showTime", api.ApiController{}.ShowTime)
	}
}
