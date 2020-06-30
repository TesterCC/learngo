package main

import (
	"fmt"
	"regexp"
)

// 14-2  https://www.bilibili.com/video/BV18Q4y1M7NV?p=54     19:33
const text = "My email is testerlyx@test123.com@eee"

const text3 = `My email is testerlyx@foxmail.com@eee, Her email is test@test.com
He email is tes23@abc.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile("testerlyx@foxmail.com")   // 确定匹配内容没错用这个，精准匹配
	//re, err := regexp.Compile("testerlyx@foxmail.com")
	match := re.FindString(text)
	fmt.Println(match)

	re2 := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)  // ``中不发生转义
	match2 := re2.FindString(text)
	fmt.Println(match2)

	match3 := re2.FindAllString(text3, -1)    //获得所有匹配
	fmt.Println(match3)

	re4 := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match4 := re4.FindAllString(text3, -1)    //获得所有匹配
	fmt.Println(match4)
}
