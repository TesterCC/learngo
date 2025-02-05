package main

import (
	"fmt"
	"time"
)

// BV19X4y1t77B
func main() {
	handler(3,4)
}

// 缺点，容易忘记或漏掉，代码比较冗余
func handler(a,b int) string{
	// get current time
	t0 := time.Now()

	if a>b {
		time.Sleep(100 * time.Millisecond)
		// computed cost time
		fmt.Printf("1. use time %d ms\n", time.Since(t0).Milliseconds())
        return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("2. use time %d ms\n", time.Since(t0).Milliseconds())
		return "ok"
	}

}
