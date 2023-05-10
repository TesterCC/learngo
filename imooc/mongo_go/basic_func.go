package main

import (
	"fmt"
	"time"
)
// 3-1
// 3-2 todo https://www.imooc.com/video/24258
func main() {
	timeDuration := 7
	fmt.Printf("倒计时：%d 秒\n", timeDuration)
	endTime := countDown(timeDuration)
	fmt.Println(endTime)
}

func countDown(timeDuration int) (endOutput string) {
	for i := timeDuration; i > 0; {
		fmt.Println(i)
		//fmt.Printf("倒计时中：%ds\n", i)
		time.Sleep(1 * time.Second)
		i--
	}

	endOutput = "倒计时完成..."
	return

}
