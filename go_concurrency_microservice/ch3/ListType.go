package main

import (
	"container/list"
	"fmt"
)

/*
P55 Go语言的列表通过双向链表方式实现，列表没有限制其内保存成员的类型
列表每次插入操作都会返回一个*list.Element结构，用以指向当前插入值所在的节点
*/

func main() {
	tmpList := list.New()

	for i := 1; i <= 10; i++ {
		tmpList.PushBack(i)
	}

	first := tmpList.PushFront(0)
	tmpList.Remove(first)

	// 配合Front()函数获取列表的头元素
	for l := tmpList.Front(); l != nil; l = l.Next() {
		fmt.Print(l.Value, " ")
	}

}
