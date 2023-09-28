package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

type Article struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func (con ApiController) Index(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "success",
	})
}

func (con ApiController) AddArticle(c *gin.Context) {
	article := &Article{}
	if err := c.ShouldBind(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, article)
	}
}
