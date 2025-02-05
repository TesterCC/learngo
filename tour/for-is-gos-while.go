package main

import "fmt"

/*

for 是 Go 中的 “while”
此时你可以去掉分号，因为 C的while 在Go中叫做for。

*/

func main() {

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
