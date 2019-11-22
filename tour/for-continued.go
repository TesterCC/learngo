package main

import "fmt"

/*
The init and post statements are optional.
初始化语句和后置语句是可选的。

for loop:
the init statement;the condition expression;the post statement {}
*/

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)    // 1024
}
