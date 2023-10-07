package main

import (
	"GoDemo1/routers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	//一定要在路由分组之前加入，否则无法读取到session
	//配置session中间件
	//创建基于cookie的存储引擎，secret参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	//设置session中间件，参数mysession根据自己的需要设置，这里使用默认值，store是前面创建的存储引擎
	r.Use(sessions.Sessions("mysession", store)) //路由分组
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
