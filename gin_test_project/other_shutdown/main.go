package main

import "github.com/gin-gonic/gin"

// 2-19 gin基础拓展:优雅关停服务器

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		
	})
}
