package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()

	// write your code
	time.Sleep(2 * time.Millisecond)  // if comment here, elapsed may return 0, because it's too fast.

	elapsed := time.Since(t0)   // 函数返回的时间间隔（elapsed）的精度是纳秒级别的

	// 精确到纳秒
	fmt.Printf("Took %dns\n", elapsed.Nanoseconds())

	// 精确到秒
	fmt.Printf("Took %.2fs\n", elapsed.Seconds())

	// default
	fmt.Printf("Took %s", elapsed)
}
