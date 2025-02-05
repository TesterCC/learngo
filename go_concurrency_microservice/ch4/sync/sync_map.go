package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
P98-99 sync.Map 并发安全字典（Go 1.9+），即添加了同步控制的字典，因为原生Map并不是并发安全的

4-8 使用sync.Map并发添加数据
*/

var syncMap sync.Map
var waitGroup sync.WaitGroup

func main() {
	routineSize := 5
	// 让主线程等待数据添加完毕
	waitGroup.Add(routineSize)
	// 并发添加数据
	for i := 0; i<routineSize; i++ {
		go addNumber(i * 10)
	}

	// 开始等待
	waitGroup.Wait()
	var size int
	// 统计数量     // syncMap.Range 无序遍历 map
	syncMap.Range(func(key, value interface{}) bool {
		size++
		fmt.Println("key-value is "+ strconv.Itoa(size)) // 辅助理解
		return true
	})

	fmt.Println("syncMap current size is " + strconv.Itoa(size))

	// 获取键为0的值     // syncMap.Load 根据 key 获取存储值
	value, ok := syncMap.Load(0); if ok{
		fmt.Println("key 0 has value", value, " ")
	}
}


func addNumber(begin int) {
	// 往syncMap中放入数据
	for i:=begin; i<begin+3; i++{
		syncMap.Store(i, i)  // 设置 key-value 对
	}

	// 通知数据已添加完毕
	waitGroup.Done()

}
