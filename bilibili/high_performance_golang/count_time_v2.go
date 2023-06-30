package main

import (
	"fmt"
	"time"
)

// BV19X4y1t77B
func main() {
	handler(3,4)
}

// 优化，使用defer，临退出时再做最后的时间打印
func handler(a,b int) string{
	// get current time
	t0 := time.Now()

	// computed cost time when exit
	// defer注册带参数函数时就会确定参数值， 如果直接 defer fmt.Printf("[D] cost time %d ms\n", time.Since(t0).Milliseconds())就会记录为注册时的时间
	defer func(){
		fmt.Printf("[D] cost time %d ms\n", time.Since(t0).Milliseconds())
	}()

	if a>b {
		time.Sleep(100 * time.Millisecond)

        return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("2. use time %d ms\n", time.Since(t0).Milliseconds())
		return "ok"
	}

}
