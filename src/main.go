package main

import "github.com/gin-gonic/gin"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "This is news page.")
	})
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
	r.Run(":8080")
}
