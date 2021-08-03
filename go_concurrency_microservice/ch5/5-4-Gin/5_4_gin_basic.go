package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// curl http://127.0.0.1:8000/ping

// curl http://127.0.0.1:8000/user/testercc

// curl http://127.0.0.1:8000/user/testercc/push


func main() {
	router := gin.Default() // 创建一个路由handler

	// 通过HTTP方法绑定路由规则和路由函数。  Gin把request和response都封装到gin.Context的上下文环境
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.String(http.StatusOK, "Hello %s", name)

	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		msg := name + " is " + action
		c.String(http.StatusOK, msg)

	})

	err := router.Run(":8000") // 监听 8000, 默认监听 8080
	if err != nil {
		return
	}
}
