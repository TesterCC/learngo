package main

import "github.com/gin-gonic/gin"

//Gin入门实战  https://www.imooc.com/learn/1175

func main() {
	r:=gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messge":"pong",
		})

	})
	r.Run() // default 8080
	
}
