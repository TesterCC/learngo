package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ServerConfig struct {
	ip   string
	port int
}

type Config struct {
	InitTime     string
	Cm           string
	ServerConfig ServerConfig
	Other        []string
}

func main() {
	// set json file name and path
	jsonPath := "./config.json"

	// fill Config data, need define struct before
	// https://www.jianshu.com/p/4d1241b02b90
	/*  序列化
		编码分为两步：

		1.创建一个新的 json 文件；
		2.将数据结构中的内容按格式写入 json 文件。
	*/

	// go 写入的 json，key 与 value 是和我们的写入数据结构体中的成员名与其值相对应的。
	info := []Config{
		{"04:00", "192.168.100.101", ServerConfig{"192.168.100.102", 8080}, []string{}},
		{"14:00", "192.168.101.101", ServerConfig{"192.168.101.102", 8080}, []string{"string-test-1", "string-test-2"}},
	}

	// create a new json file
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		fmt.Println("file create failed", err.Error())
		return
	}

	defer jsonFile.Close()

	// create json encoder
	encoder := json.NewEncoder(jsonFile)

	// encode
	err = encoder.Encode(info)

	if err != nil {
		fmt.Println("Encode Error", err.Error())
		log.Printf("encode err [ %v ]", err)
		return
	} else {
		fmt.Println("Encode Success")
	}
}
