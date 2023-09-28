package main

import (
	"GoDemo1/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//路由分组
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "success",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		user := User{
			Name: "Tom",
			Age:  18,
		}
		c.JSON(200, user)
	})

	r.GET("/news1", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "这是后端传入的数据",
		})
	})

	r.GET("/news2", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"page": page,
		})
	})

	r.GET("/getuser", func(c *gin.Context) {
		user := &UserInfo{}
		//fmt.Println(*user)
		if err := c.ShouldBind(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, user)
			//fmt.Println(*user)
		}
	})

	//动态路由
	r.GET("list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "动态路由的值为：%v", cid)
	})

	r.Run(":8080")
}
