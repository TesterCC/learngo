package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

/*
P109-110 5-2 Go Web请求体解析
login.tpl  登录页表单

增加 /login路由，通过登录页表单提交信息，再服务端进行验证登录的结果
*/

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[+] method:", r.Method)   // 获取请求的方法并打印
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.tpl")
		log.Println(t.Execute(w,nil))
	} else if r.Method== "POST" {
		// 请求的是登录数据，那么执行登录的逻辑判断
		_ = r.ParseForm()
		fmt.Println("[+] username:",r.Form["username"])
		fmt.Println("[+] password:",r.Form["password"])   // 例子为了方便用明文，实际绝对不能用明文

		// 验证密码是否正确
        if pwd := r.Form.Get("password"); pwd == "1234567"{
        	fmt.Fprintf(w, "Welcome to login, Hello %s !", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "Password Error, please check!")
		}


	} else {
		log.Println("Don't support the method.")
	}

}

func sayHelloName(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHelloName)       // 设置访问路由
	http.HandleFunc("/login", login)         // 设置访问路由
	err := http.ListenAndServe(":8080", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
