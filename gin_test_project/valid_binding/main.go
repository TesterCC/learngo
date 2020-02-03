package main

import (
	"github.com/gin-gonic/gin"
)

// 2-10 to 2-11 gin基础：验证请求参数 - 结构体验证

/*
验证请求参数：
1.结构体验证
2.自定义验证
3.升级验证 - 支持多语言错误信息  (根据请求的语言进行错误信息转化)


Test:
curl -X GET "http://127.0.0.1:8888/testing?age=18&name=Lily&address=beijing"


gin_scaffold课程代码：
gin入门实战 - 基础精髓: https://github.com/e421083458/hello_gin
后端代码：https://github.com/e421083458/gin_scaffold
前端代码：https://github.com/e421083458/vue-admin
*/

type Person struct {
	// Gin具体默认封装规则：https://godoc.org/gopkg.in/go-playground/validator.v8
	Age     int    `form:"age" binding:"required,gt=10"` //  验证规则要同时满足时用","分隔； 任何一个能满足用"|"分隔
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", func(c *gin.Context) {
		// 参数已经绑定到结构体，再基于这个结构体进行验证
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}

		c.String(200, "%v", person)

	})

	r.Run(":8888")
}
