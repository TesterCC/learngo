package main

import "github.com/gin-gonic/gin"

// https://gin-gonic.com/zh-cn/docs/testing/

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}