package main

import (
	"fmt"
	"time"
)

// go timestamp

func main() {
	//after go version 1.17, go support millisecond (毫秒，千分之一秒) and microsecond (微秒，百万分之一秒)

	now := time.Now()  // return current localtime
	fmt.Println(now.Unix())            // second, 10-digital,      e.g.: 1667120289
	fmt.Println(now.UnixMilli())       // millisecond, 13-digital, e.g.: 1667120289071
	fmt.Println(now.UnixMicro())       // microsecond, 16-digital, e.g.: 1667120289071470
	fmt.Println(now.UnixNano())        // nanosecond, 19-digital,  e.g.: 1667120289071470000     (纳秒,十亿分之一秒；)

}