package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) GetCookie(c *gin.Context) {
	c.SetCookie("username", "Xiaobo", 3600, "/", "localhost", false, true)
}

func (con IndexController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "Xiaobo", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete cookie success",
	})
}

func (con IndexController) News(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	page := c.DefaultQuery("page", "1")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
		"page": page,
	})
}

func (con IndexController) News2(c *gin.Context) {
	usernanme, err := c.Cookie("username")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username": usernanme,
		})
	}
}
