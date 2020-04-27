package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 2-15 to 2-16 gin基础：中间件-使用

func main() {
	//将日志打印到文件中
	f,_ := os.Create("gin.log")   // gin.log在这个go文件的上一级目录
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)  // 将错误日志也重定向到文件中

	//r:=gin.Default()   // 默认实现2个中间件 Logger(), Recovery()
	r:=gin.New()
	//r.Use(gin.Logger())   // 此时触发panic会让服务挂掉
	r.Use(gin.Logger(), gin.Recovery())  // Recovery()使得发生错误时服务不会直接挂掉
	r.GET("/test", func(c *gin.Context) {
		panic("test panic")
		name:=c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})
	r.Run()
}
