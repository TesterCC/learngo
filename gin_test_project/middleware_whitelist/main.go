package main

import "github.com/gin-gonic/gin"

/*
2-17 gin基础：中间件-自定义中间件
*/

// 实现自定义中间件
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",   // 实际建议存在DB，每次使用时再获取
		}
		flag := false

		clientIP := c.ClientIP()   // 访问者IP

		for _,host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(401, "%s, not in IP White List", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	// 期待的是r.Use(IPAuthMiddleware())  实现参考 gin.Logger(), gin.Recovery() 的代码

	r.Use(IPAuthMiddleware())
	r.GET("test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	r.Run()
}
