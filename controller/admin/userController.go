package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 定义为结构体，可以实现继承
type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "This is user page.")
}

func (con UserController) Add(c *gin.Context) {
	userinfo := &UserInfo{}
	if err := c.ShouldBind(userinfo); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, userinfo)
	}
}

func EditUser(c *gin.Context) {
	userinfo := &UserInfo{}
	if err := c.ShouldBind(userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, userinfo)
	}
}
