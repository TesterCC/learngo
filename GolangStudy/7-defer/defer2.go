package main

import "fmt"

/*
知识点二：defer和return谁先谁后

return func called...
defer func called...

return先执行，defer后执行

defer是当前程序逻辑结束后才出栈
*/

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
