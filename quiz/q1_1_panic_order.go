package main

import "fmt"

func main() {
	defer_call2()
}

func defer_call2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=", err)
		}
	}()
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

}

/*
https://my.oschina.net/qiangmzsx/blog/1478739
打印后
打印中
打印前
one= 触发异常

为什么没有加recover()时候，panic执行顺序不定呢？
defer的执行顺序肯定是FILO的，但是没有被recover的panic协程（线程）可能争夺CPU的顺序比defer快，
所以造成了这样的情况，也可能是写缓存问题，所以对panic进行recover将其加入到defer队列中
*/
