package main

import "github.com/gin-gonic/gin"

/*
2-7 gin基础：请求路由 - 泛绑定    https://www.imooc.com/video/20183
泛绑定前缀
通过前缀匹配，所有匹配都会打到一个相同的方法上
*/

func main() {
	r := gin.Default()
	// 所有user前缀的都会打到回复hello world的这个地方
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world" )
	})
    r.Run()
}


/* testcase
curl -X GET "http://127.0.0.1:8080/user/"
curl -X GET "http://127.0.0.1:8080/user/e"
curl -X GET "http://127.0.0.1:8080/user/1111"
*/