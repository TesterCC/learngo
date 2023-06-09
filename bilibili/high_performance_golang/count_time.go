package main

import (
	"fmt"
	"time"
)

// BV19X4y1t77B
func main() {
	handler(3,4)
}

func handler(a,b int) string{
	t0 := time.Now()

	if a>b {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("1. use time %d ms\n", time.Since(t0).Milliseconds())
        return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("2. use time %d ms\n", time.Since(t0).Milliseconds())
		return "ok"
	}

}
