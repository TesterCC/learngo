package main

import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/autotls"
/*
2-21 gin基础拓展:自动化证书配置

http://www.github.com/gin-gonic/autotls
*/

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test tls")
	})

	//调用自动下载证书的包即可
	autotls.Run(r, "xxx.com")
	// 测试，请求 xxx.com/test
}
