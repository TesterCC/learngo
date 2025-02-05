package main

import "github.com/gin-gonic/gin"

// [2-3 gin基础：请求路由 - 多种请求 video](https://www.imooc.com/video/20179)

/*
test command:

curl -X GET "http://127.0.0.1:8080/get"
curl -X POST "http://127.0.0.1:8080/post"
curl -X DELETE "http://127.0.0.1:8080/delete"

curl -X PUT "http://127.0.0.1:8080/any"
curl -X CONNECT "http://127.0.0.1:8080/any"
*/

func main() {
	r:=gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(200,"get method")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200,"post method")
	})

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete method")
	})

	// 如果一个请求路由想对应多个请求方法
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
}
