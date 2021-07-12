package main

import "net/http"

/*
P109-110 5-2 Go Web请求体解析
login.tpl  登录页表单

增加 /login路由，通过登录页表单提交信息，再服务端进行验证登录的结果
*/

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHelloName)
}

func sayHelloName(writer http.ResponseWriter, request *http.Request) {
	
}