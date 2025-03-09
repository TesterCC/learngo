package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

/*
ref: https://www.bilibili.com/video/BV1HJPVeGE5h
GIN vs Fiber 编程与性能对比
*/

func InitGin() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return engine
}

func InitFiber() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	return app
}

// bind and validation, gin use binding, fiber use validate
type Arg struct {
	Name string `form:"name" binding:"required, gt=4" validate:"required,gt=4"`
	Age  int    `form:"age" binding:"required,gt=0" validate:"required,gt=0"`
}

func GinHandler(ctx *gin.Context) {
	var arg Arg

	ctx.ShouldBind(&arg)

	//if err := ctx.ShouldBind(&arg); err != nil { // binding args and validation
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	ctx.Header("company", "sec")                                         // response header
	ctx.SetCookie("name", "value", 3600, "/", "localhost", false, false) // set cookie
	ctx.JSON(http.StatusOK, arg)                                         // response body                                              // response body
}

func FiberHandler(ctx fiber.Ctx) error {
	var arg Arg
	ctx.Bind().Form(&arg)     // 绑定参数并校验
	ctx.Set("company", "sec") // 响应头

	ctx.Cookie(&fiber.Cookie{Name: "name", Value: "value", Path: "/", MaxAge: 3600, Domain: "localhost", Secure: false, HTTPOnly: false})

	ctx.JSON(arg) // 响应体
	return nil
}

func GinServe(port int) {
	engine := InitGin()
	engine.POST("/", GinHandler)
	engine.Run(":" + strconv.Itoa(port))
}

func FiberServe(port int) {
	app := InitFiber()
	app.Post("/", FiberHandler)
	app.Listen(":" + strconv.Itoa(port))
}

//func Post(uri string) {
//	resp, _ := http.PostForm(uri, url.Values{
//		"name": {"张三"},
//		"age":  {"18"},
//	})
//
//	for k, v := range resp.Header {
//		fmt.Println("%s=%s\n", k, v[0])
//	}
//}

// 差一段函数

/*
go test ./web -bench=^BenchmarkGin$ -run=^$ -count=1
go test ./web -bench=^BenchmarkFiber$ -run=^$ -count=1
*/

func BenchmarkGin(b *testing.B) {
	port := 5678
	go GinServe(port) // 携程启动 launch Server
	time.Sleep(3 * time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		http.PostForm(fmt.Sprintf("http://localhost:%d/", port), url.Values{"name": {"张三"},
			"age": {"18"}})

	}
}

func BenchmarkFiber(b *testing.B) {
	port := 5678
	go FiberServe(port)
	time.Sleep(3 * time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		http.PostForm(fmt.Sprintf("http://localhost:%d/", port), url.Values{"name": {"张三"},
			"age": {"18"}})
	}
}
