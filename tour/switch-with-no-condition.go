package main

import (
	"fmt"
	"time"
)

/*
https://tour.go-zh.org/flowcontrol/11
https://go.dev/tour/flowcontrol/11

无条件 switch

无条件的 switch 同 switch true 一样。
这种形式能将一长串 if-then-else 写得更加清晰。

格式化字符串：
    使用 %v 打印默认格式。

    使用 %+v 打印详细格式（适合调试）。

    使用 %#v 打印 Go 语法格式。

    结合 Format 方法自定义格式，用 %s 打印。

    使用 %d 打印 Unix 时间戳或时间部分。
*/

func main() {
	t := time.Now()

	// format print
	fmt.Printf("Default format: %v\n", t)                                        // default directly print
	fmt.Printf("Detailed format: %+v\n", t)                                      // 会显示更多细节。
	fmt.Printf("Go syntax format: %#v\n", t)                                     // 会打印值的 Go 语法表示形式。
	fmt.Printf("【Attetion】Custom format: %s\n", t.Format("2006-01-02 15:04:05")) // 美式风格：月日时分秒年 使用 Format 方法将时间格式化为字符串，然后用 %s 打印。
	fmt.Printf("Unix timestamp (seconds): %d\n", t.Unix())
	fmt.Printf("Unix timestamp (nanoseconds): %d\n", t.UnixNano())

	fmt.Printf("Year: %d\n", t.Year())
	fmt.Printf("Month: %d\n", t.Month())
	fmt.Printf("Day: %d\n", t.Day())
	fmt.Printf("Hour: %d\n", t.Hour()) // local use system time
	fmt.Printf("Minute: %d\n", t.Minute())
	fmt.Printf("Second: %d\n", t.Second())

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
