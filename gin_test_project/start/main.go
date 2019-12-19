package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
ref:
https://www.cnblogs.com/ningskyer/articles/6297356.html

Learn Basic Gin
*/

func main() {
	router:=gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})


	//router.Run()   // method 1, default 8080
    //http.ListenAndServe(":8888", router)   // method 2

	//custom http config   method 3
	s := &http.Server{
		Addr:		":9999",
		Handler: 	router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:	   10 * time.Second,
		MaxHeaderBytes: 	1 << 20,
	}

	s.ListenAndServe()
	// test:  curl -X GET "http://127.0.0.1:8080/ping"

}