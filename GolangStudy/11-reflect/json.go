package main

import (
	"encoding/json"
	"fmt"
)

/*
25-结构体标签在json中的应用 (这个在实际开发中高频使用: 1.json编解码 2.orm映射关系)
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=25
*/

type Movie struct {
	Title  string   `json:"title"` // 不写tag也可以转，默认是以属性名称为key
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

// 把结构体转化为json
func main() {
	movie := Movie{"社交网络", 2010, 10, []string{"Alice", "Bob"}}

	// 1. Encode编码的过程：struct --> json str  序列化： 对象 -> 字节序列
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	// jsonStr = {"title":"社交网络","year":2010,"rmb":10,"actors":["Alice","Bob"]} // 可以把输出的json str放到在线json校验网站进行校验 https://c.runoob.com/front-end/53/

	// 2. Decode解码的过程 json str --> struct  反序列化：字节序列 -> 对象
	myMovice := Movie{}
	err = json.Unmarshal(jsonStr, &myMovice)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
		return
	}

	fmt.Printf("%T  %v", myMovice, myMovice)
}
