package main

import "fmt"

/*
https://my.oschina.net/qiangmzsx/blog/1478739
1.写出下面代码输出内容。

考点：defer执行顺序

解答：
defer 是后进先出。 FILO
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
*/

func main() {
	defer_call()
}


func defer_call() {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()

	panic("触发异常")
}

/*
打印后
打印中
打印前
panic: 触发异常
*/