package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// https://www.jianshu.com/p/4d1241b02b90

/*
    反序列化
	解码分为两步：

	1.打开待解码的 json 文件；
	2.使用 json 包提供的方法解码 json 文件到数据结构中。
*/

// Go 语言中的 json 包会自动识别 json 项到相应的成员中（json 项与成员变量名对应，不区分大小写）
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
	jsonPath := "./config.json"

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Printf("open json file %v error [ %v ]", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	var conf []Config
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("decode error [ %v ]", err)
		return
	} else {
		fmt.Println(conf)
		//fmt.Println(conf[0])
		//fmt.Println(conf[0].Cm)
		//fmt.Println(conf[0].InitTime)
		//fmt.Println(conf[0].ServerConfig)

		for index, item := range conf {
			fmt.Println(index, "-->", item)
			fmt.Println("Cm: ", item.Cm)
			fmt.Println("InitTime: ", item.InitTime)
			fmt.Println("ServerConfig: ", item.ServerConfig)
			fmt.Println("Other: ", item.Other)
		}
	}
}
