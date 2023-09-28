package main

import (
	"GoDemo1/routers"
	"fmt"
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

// 路由中间件，Next()之后的执行顺序是从后向前。
func initMiddlewareOne(c *gin.Context) {
	fmt.Println("1-这是中间件-initMiddlewareOne")
	c.Next()
	fmt.Println("3-这是中间件-initMiddlewareOne")
}

func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1-这是中间件-initMiddlewareTwo")
	c.Next()
	fmt.Println("3-这是中间件-initMiddlewareTwo")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//路由分组
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	//配置全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)

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
	//中间件
	//r.GET("/news3", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
	//	name := c.Query("name")
	//	age := c.Query("age")
	//	page := c.DefaultQuery("page", "1")
	//	fmt.Println("2-这是路由函数")
	//	c.JSON(http.StatusOK, gin.H{
	//		"name": name,
	//		"age":  age,
	//		"page": page,
	//	})
	//})

	//动态路由
	r.GET("list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "动态路由的值为：%v", cid)
	})

	r.Run(":8080")
}
