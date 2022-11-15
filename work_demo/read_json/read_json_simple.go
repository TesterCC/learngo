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

// 只有一组 json 项
// Go 语言中的 json 包会自动识别 json 项到相应的成员中（json 项与成员变量名对应，不区分大小写）
type Config struct {
	Repo string
	Name string
	Blog string
}


func main() {
	jsonPath := "./info.json"

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Printf("open json file %v error [ %v ]", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	var conf Config
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("decode error [ %v ]", err)
		return
	} else {
		fmt.Println(conf)
		fmt.Println(conf.Name)
		fmt.Println(conf.Blog)
		fmt.Println(conf.Repo)
	}
}