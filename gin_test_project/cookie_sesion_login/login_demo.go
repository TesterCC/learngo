package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ref: https://blog.csdn.net/qq_55752792/article/details/126522565

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("login"); err == nil {
			if cookie == "yes" {
				c.Next()
			}
		} else {
			// 返回错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "没有登录"})
			// 若验证不通过，不再调用后续的函数处理
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("login", "yes", 60, "/", "", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "登陆成功，能访问home"})
	})

	r.GET("/logout", func(c *gin.Context) {
		// clear cookie
		c.SetCookie("login","",-1,"/","", false, true)
		c.String(200, "Logout success!")
	})
	r.Run(":8080")
}
