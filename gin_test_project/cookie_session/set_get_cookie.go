package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
Gin 中使用 Cookie,Session
https://blog.csdn.net/qq_55752792/article/details/126522565

Cookie介绍和使用

HTTP 是无状态协议。简单地说，当你浏览了一个页面，然后转到同一个网站的另一个页 面，服务器无法认识到这是同一个浏览器在访问同一个网站。
每一次的访问，都是没有任何 关系的。如果我们要实现多个页面之间共享数据的话我们就可以使用 Cookie，Session或Token实 现

cookie 是存储浏览器中的键值对，可以让我们用同一个浏览器访问同一个域名 的时候共享数据。
*/

func main() {
	r := gin.Default()
	r.GET("set_cookie", func(c *gin.Context) {
		c.SetCookie("name", "lxx", 60, "/", "", false, true)
		c.SetCookie("age", "19", 600, "", "", false, true)
		c.String(200, "cookie设置成功")
	})
	r.GET("get_cookie", func(c *gin.Context) {
		name, err := c.Cookie("name")
		if err != nil {
			c.String(200, "cookie获取打印失败,错误是:%s", err)
			return
		}
		fmt.Println("name的cookie值为：", name)
		c.String(200, "cookie获取打印成功")

	})
	r.Run(":8080")
}
