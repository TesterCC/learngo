package main

import "fmt"

/*
31-channel与range
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=31
*/

// 2个goroutine通信，尝试关闭channel
func main() {
	// 创建无缓冲channel
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 将i写入channel
		}

		// close可以关闭一个channel   这里 sub goroutine遍历完才会关闭channel
		close(c)
	}()

	// main goroutine
	/*
		for {
			// ok若为true表示channel没有关闭，若为false表示channel已经关闭
			if data,ok := <-c; ok {
				fmt.Println(data)
			} else {
				break  // 跳出当前的死循环
			}
		}
	*/
	// 用range简写上面那段代码
	// 可以使用range来迭代不断操作channel
	// 当c中有数据，range返回；当c中无数据且未关闭，range就会阻塞
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished ...")
}
