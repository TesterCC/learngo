package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
go高级编程 1.2.7 P21

使用标准库自带的net/http包构造了一个独立运行的http服务
*/

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345/")

	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request){
		s := fmt.Sprintf("您好，世界！-- Time: %s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("%v\n", s)
	})

	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
