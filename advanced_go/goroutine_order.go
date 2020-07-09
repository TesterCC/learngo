package main

// P421-422 Appendix 9.不同Goroutine之间不满足顺序一致性内存模型
// 解决办法是用显示同步
// msg的写入是在channel发送之前，所以能保证打印hello, world
var msg string
//var done bool
var done = make(chan bool)

func setup() {
	msg = "hello, world"
	done <- true
}

func main() {
	go setup()
    <- done
	println(msg)
}
