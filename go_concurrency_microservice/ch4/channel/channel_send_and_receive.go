package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
P91 channel
channel 作为一个具备长度的容器，可以被遍历。
需要注意的是，在channel关闭后不允许再往通道中放入数据，不然会抛出panic
*/

func printInput(ch chan string){
	// 使用 for 循环从 channel 中读取数据
	for val := range ch {
		// 读取到结束符号
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	// 创建一个无缓冲的channel
	ch := make(chan string)
	go printInput(ch)

	// 从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val   // 表示 val 将被发送到 channel中
		if val == "EOF" {
			fmt.Println("End the game!")
			break
		}
	}
	// 程序最后关闭ch
	defer close(ch)
}