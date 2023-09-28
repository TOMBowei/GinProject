package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
	//中间件设置数据
	c.Set("username", "xiaobo")
	c.Next()
	fmt.Println("这是中间件")

	//	定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
