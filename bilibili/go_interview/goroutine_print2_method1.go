package main

import (
	"fmt"
	"sync"
)

/*
https://www.nowcoder.com/discuss/712384113688199168
标准解法是1，还有两种解法

两个协程交替打印  sxf手撕
题目描述：使用两个go协程，交替打印26个字母和数字，输出结果: A1B2...Z26。
题目解析：解题关键在于当我们开了两个协程时，如何控制协程1打印一个字母后，暂停执行，通知到协程2打印数字。当协程2打印完数字后，暂停执行，反过来通知到协程1打印下一个字母。
*/

/*
方法一：
解析：首先定义两个无缓冲的channel: ch1和ch2。
开两个go协程分别循环打印26个字母和数字。在循环体里，每次打印之前，都从无缓冲的channel里读取数据。
在未给ch1和ch2传入数据时，两个协程for循环的第一次执行，都会由于无法从无缓冲的channel里获取数据而阻塞。
当我们在主函数里执行`ch1 <- struct{}{}`，协程1的阻塞得以释放，顺利打印出第一个字母。
打印完成后，执行`ch2 <- struct{}{}`,此时下一个for循环由于ch1没有数据被阻塞。
而协程2由于监听到ch2传入数据，停止阻塞并打印第一个数字。
打印完毕后，执行`ch1 <- struct{}{}`,下一次for循环由于ch2没有数据被阻塞，却又释放了协程1，打出第二个字母。
以此类推，两个协程循环往复地交替打印着。
当打印完最后一个数字26时，ch2不再有协程收发数据(两个for循环都执行完成)，但ch1会最后存入一次数据，所以要在协程1最后，从ch1中最后接收一次数据，避免程序死锁。
*/

// 协程1
func printLetters(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	// 在子协程函数的开头使用 defer 关键字调用。
	// wg.Done() 方法用于通知 WaitGroup 一个协程已经完成了它的任务。
	// 使用 defer 关键字可以确保无论协程函数是正常返回还是发生异常，wg.Done() 都会被调用。
	// 这样可以保证 WaitGroup 的计数器正确减少。
	// 如果不使用 defer，在某些异常情况下可能会导致 wg.Done() 没有被调用，从而使 WaitGroup 的计数器无法正确减少，最终导致 wg.Wait() 一直阻塞。
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i++ {
		<-ch1
		fmt.Printf("%c", i)
		ch2 <- struct{}{}
	}
	<-ch1
}

// 协程2
func printNumbers(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 26; i++ {
		<-ch2
		fmt.Printf("%d", i)
		ch1 <- struct{}{}
	}

}

func main() {

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// 协调多个协程（goroutine）的执行，确保主协程能够等待所有子协程都完成任务之后再继续执行或者退出。
	var wg sync.WaitGroup
	// 在启动子协程之前调用。wg.Add(n) 方法用于设置需要等待的协程数量。
	// 确保 WaitGroup 知道需要等待两个协程完成任务，这样后续的 wg.Wait() 才能正确工作。
	// 如果在启动协程之后调用 wg.Add，可能会导致 WaitGroup 状态不一致，从而无法正确等待所有协程结束。
	wg.Add(2)

	go printNumbers(ch1, ch2, &wg)
	go printLetters(ch1, ch2, &wg)

	ch1 <- struct{}{}

	// 在启动所有子协程之后调用。
	// wg.Wait() 方法会阻塞当前协程（这里是主协程），直到 WaitGroup 的计数器变为 0。
	// 所以当所有子协程都完成任务后，计数器会变为 0，wg.Wait() 会解除阻塞，主协程可以继续执行后续代码。
	// 如果不调用 wg.Wait()，主协程可能会在子协程还未完成任务时就退出，导致子协程无法正常执行完毕。
	wg.Wait()

	//fmt.Println("\ntest over")  // debug
}
