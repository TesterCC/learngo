package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
https://tour.go-zh.org/flowcontrol/10
https://go.dev/tour/flowcontrol/10

switch 的求值顺序
switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
例如，

switch i {
case 0:
case f():
}
在 i==0 时，f 不会被调用。
注意： Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始， 该值的意义留给读者去发现。
p.s. 是 Go 语言的发布时间，具有纪念意义。
*/

func main() {
	fmt.Println("When's Saturday?")

	// 返回当前是星期几，类型为 time.Weekday
	today := time.Now().Weekday()

	fmt.Println("type: ", reflect.TypeOf(today))
	fmt.Println("today:", today)

	// time.Saturday 6
	// time.Weekday 是一个整数类型（int），表示一周中的某一天，范围是 0（周日）到 6（周六）。
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}
}
