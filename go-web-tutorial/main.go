package main

import "net/http"

func main() {
	//1.注册一个函数，使其可以对web请求进行响应
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	//2.启动http server
	http.ListenAndServe("localhost:8080", nil) // DefaultServeMux
}
