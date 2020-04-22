package main

import "github.com/gin-gonic/gin"

// Task: gin框架中结合gorilla实现webSocket
// gin框架写rest接口很方便且性能好
// ref: https://blog.csdn.net/lihao19910921/article/details/81907742

func main() {
	bindAddress := "localhost:3030"

	r := gin.Default()

	// 用路径/test 监听 get 请求
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, []string{"ab","abc","135","246"})
	})

	r.Run(bindAddress)
}
