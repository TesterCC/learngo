package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 2-19 gin基础拓展:优雅关停服务器, 看gin启动对比图

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(5*time.Second)
		c.String(200, "Hello test in April")
	})

	srv:=&http.Server{
		Addr:":8085",
		Handler:r,
	}

	go func() {
		if err:=srv.ListenAndServe(); err!= nil && err!=http.ErrServerClosed{
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//真正的请求拦截
	quit:=make(chan os.Signal)  //定义channel
	signal.Notify(quit,syscall.SIGINT, syscall.SIGTERM)
	<-quit  //阻塞channel
	log.Println("Shutdown Elegant Server...")
 	//创建超时的上下文
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	//真正关闭server
	if err:=srv.Shutdown(ctx); err!=nil {
		log.Fatal("server shutdown occurs error: ", err)
	}
	log.Println("Server Exiting")

}
