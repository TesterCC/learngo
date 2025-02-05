package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// [2-4 gin基础： 请求路由 - 静态文件夹](https://www.imooc.com/video/20180)
func main() {
	r:=gin.Default()

	// method 1 先设置路由，再设置相对文件夹
	r.Static("/assets","./assets")

	// method 2 先设置路由，再设置资源
	r.StaticFS("/static", http.Dir("static"))

	// method 3 设置单个静态文件,先设置路由，再设置资源
	r.StaticFile("/favicon.ico", "./favicon.ico")

    r.Run()
}

// 命令行运行，否则会找不到static文件夹
// cd ~/gin_test_project/router_static
// go build -o router_static && ./router_static


/*  测试3中静态路由设置
curl http://127.0.0.1:8080/assets/div_demo.html
curl http://127.0.0.1:8080/static/div_float.html
curl http://127.0.0.1:8080/favicon.ico
*/