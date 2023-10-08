package admin

import (
	"GoDemo1/models"
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

func (con UserController) ShowUser(c *gin.Context) {
	//得到的是空接口类型，需要使用类型断言。
	username, err := c.Get("username")
	if err != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	}
}

// gorm相关
func (con UserController) Adduser(c *gin.Context) {
	//user := &models.User{
	//	Username: "丁丁",
	//	Age:      2,
	//	Email:    "123",
	//	AddTime:  1234567890,
	//}
	//models.DB.Create(&user)

	user := &models.User{
		Username: "啦啦",
		Age:      20,
		Email:    "123",
		AddTime:  1,
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func (con UserController) Edituser(c *gin.Context) {
	////保存所有数据
	//user := models.User{Id: 1}
	//models.DB.Find(&user)
	//user.Age = 18
	//models.DB.Save(&user)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code":   200,
	//	"msg":    "修改成功",
	//	"result": user,
	//})
	//	更新单个字段
	user := models.User{}
	models.DB.Model(&user).Where("id=?", 1).Update("age", 20)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

func (con UserController) Deleteuser(c *gin.Context) {
	//删除数据
	user := models.User{}
	models.DB.Where("id=?", 3).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func (con UserController) Listuser(c *gin.Context) {
	//查询数据库
	//userList := models.User{}
	//models.DB.Find(&userList)
	//c.JSON(http.StatusOK, gin.H{
	//	"code":   200,
	//	"msg":    "获取成功",
	//	"result": userList,
	//})
	//	查询年龄小于20的用户
	userList := []models.User{}
	models.DB.Where("age<?", 20).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取成功",
		"result": userList,
	})
}
