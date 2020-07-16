package main

import (
	"fmt"
	"os"
)

// 获取系统环境变量
func main() {
	conn := os.Getenv("MYSQL_CONFIG")
	if conn == "" {
		fmt.Println("conn is null, please check it")
	} else {
		fmt.Printf("%T ->  %v\n",conn,conn)
	}
}
