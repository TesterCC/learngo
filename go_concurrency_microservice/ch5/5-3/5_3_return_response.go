package main

import (
	json2 "encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string
	Habits []string
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom header")
	w.WriteHeader(201)

	// &取址   *指针
	user := &User{
		Name: "TesterCC",
		Habits: []string{"CTF","PenTest","Python","Go","Rust"},
	}
	json, _:= json2.Marshal(user)
	w.Write(json)
}


func main() {
	http.HandleFunc("/write", write)  // 设置访问的路由
	err := http.ListenAndServe(":8080", nil)  // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

