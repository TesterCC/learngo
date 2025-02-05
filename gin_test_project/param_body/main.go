package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// 2-9 gin基础：获取请求参数 - 获取body内容  https://www.imooc.com/video/20185
// get form data

// curl测试传入body
// curl -X POST "http://127.0.0.1:7777/test" -d '{"name":"get_body_param"}'
// curl -X POST "http://127.0.0.1:7777/test" -d '{"name":"get_body_param", "ENV":"Go"}'
/*
curl -X POST "http://127.0.0.1:7777/test" -d 'first_name=Lily&last_name=Brown'

[POST]Lily, Brown, first_name=Lily&last_name=Brown
*/

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort() // 把输入结束
		}

		// 测试body带特定参数的返回
		//把bodyBytes的内容回存到Body中
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// 获取一个不带默认值的参数
		firstName := c.PostForm("first_name")
		// 获取一个带默认值的参数
		lastName := c.DefaultPostForm("last_name", "last_default_name")
		// first name/given name 是名字,family name/last name 是姓氏
		c.String(http.StatusOK, "[POST]%s, %s, %s", firstName, lastName, string(bodyBytes))

		//c.String(http.StatusOK, string(bodyBytes))
	})
	r.Run(":7777")
}
