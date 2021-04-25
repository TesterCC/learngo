package main

import (
	"github.com/kataras/iris/v12"
)

// curl http://localhost:8989/test

func main() {
	app := iris.New()

	// simple demo
	//app.Get("/", func(context context.Context) {})
	//app.Run(iris.Addr(":9999"))

	// ref:  https://www.bilibili.com/video/BV11E411w7Br
	ret := map[string]string{}
	ret["result"] = "test"
	app.Handle("GET", "/test", func(context iris.Context) {
		context.JSON(ret)
	})

	app.Run(iris.Addr(":8989"))
}
