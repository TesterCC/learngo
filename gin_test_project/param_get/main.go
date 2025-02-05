package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
2-8 gin基础：获取请求参数 - 获取GET参数   POST同理

testcase:
curl -X GET "http://127.0.0.1:8888/test"
curl -X GET "http://127.0.0.1:8888/test?first_name=Lily"
curl -X GET "http://127.0.0.1:8888/test?first_name=Ada&last_name=Lovelace"

curl -X POST "http://127.0.0.1:8888/test"
curl -X POST "http://127.0.0.1:8888/test?first_name=Ada&last_name=Lovelace"
*/

func main() {
	r := gin.Default()

	//获取GET请求参数
	r.GET("/test", func(c *gin.Context) {
		// 获取一个不带默认值的参数
		firstName := c.Query("first_name")
		// 获取一个带默认值的参数
		lastName := c.DefaultQuery("last_name", "last_default_name")
		// first name/given name 是名字,family name/last name 是姓氏
		c.String(http.StatusOK, "[GET]Your name is %s, %s", firstName,lastName)

	})
    // 获取POST请求参数
	r.POST("/test", func(c *gin.Context) {
		// 获取一个不带默认值的参数
		firstName := c.Query("first_name")
		// 获取一个带默认值的参数
		lastName := c.DefaultQuery("last_name", "last_default_name")
		// first name/given name 是名字,family name/last name 是姓氏
		c.String(http.StatusOK, "[POST]Hello! %s, %s", firstName,lastName)
	})

	r.Run(":8888")  //  指定端口启动
}
