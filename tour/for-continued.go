package main

import "fmt"

/*
https://tour.go-zh.org/flowcontrol/2
https://go.dev/tour/flowcontrol/2

The init and post statements are optional.
初始化语句和后置语句是可选的。

for loop:
the init statement;the condition expression;the post statement {}
*/

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Printf("...current sum: %d\n", sum) // 2 4 8 16 32 ... 1024
	}
	fmt.Println(sum) // 1024
}
