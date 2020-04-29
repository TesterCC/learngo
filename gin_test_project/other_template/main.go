package main

import "github.com/gin-gonic/gin"

/*
2-20 gin基础拓展:模板渲染

cd /learngo/gin_test_project/other_template
go run main.go

更多使用参照官网go template的文档
*/

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")   // goland中运行会找不到文件夹，在Terminal中运行
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			// gin.H{} 实际上是一个map interface.  H is a shortcut for map[string]interface{}
			"title": "index.html",
			"flag": true,
		})
	})
	r.Run()
}
