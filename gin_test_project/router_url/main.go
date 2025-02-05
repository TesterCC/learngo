package main

import "github.com/gin-gonic/gin"

// 2-5 gin基础：请求路由 - 参数作为url   https://www.imooc.com/video/20217

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			// 返回一个map
			"name": c.Param("name"),
			"id": c.Param("id"),
		})
	})

	r.Run()
	
}

// curl -X GET "http://127.0.0.1:8080/cat/10000"
// curl http://127.0.0.1:8080/testcc/10001
//{"id":"10001","name":"testcc"}