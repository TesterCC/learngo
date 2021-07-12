package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
P104-105 5.1 快速构建一个Go Web服务
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   // 3.解析参数，默认是不会解析的
	fmt.Println(r.Form) // 4.输出到服务端的打印信息
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Host: ", r.Host)

	// 4.1 打印Form表单中的键值对value
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		//fmt.Println("value: ", v)   // value:  [CC]
		fmt.Println("val: ", strings.Join(v,""))   // val:  CC
	}

	_, _ = fmt.Fprintf(w, "Hello Web, %s!", r.Form.Get("name")) // 5.写入到w的是输出到客户端的内容  POSTMAN的响应结果中可见

}

func main() {
	http.HandleFunc("/hello", sayHello)           // 1.设置访问的路由
	err := http.ListenAndServe(":8080", nil) // 2.设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
