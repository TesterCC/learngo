package main

import "fmt"

/*

https://tour.go-zh.org/moretypes/17
https://go.dev/tour/moretypes/17

可以将下标或值赋予 _ 来忽略它。

若你只需要索引，忽略第二个变量即可。
e.g.:
for i := range pow

*/

func main() {
	pow := make([]int, 10)

	// 若你只需要索引，忽略第二个变量即可。
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i    // 把1左移i位。左移操作就是最高位抛弃 后面的每一位向zd左前进一位 最后一位补0
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
