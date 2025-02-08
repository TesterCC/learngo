package main

import "fmt"

/*
https://tour.go-zh.org/flowcontrol/3
https://go.dev/tour/flowcontrol/3

for 是 Go 中的「while」
此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。

*/

func main() {

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
