package main

/*
https://tour.go-zh.org/generics/2
https://go.dev/tour/generics/2


泛型类型
除了泛型函数之外，Go 还支持泛型类型。 类型可以使用类型参数进行参数化，这对于实现通用数据结构非常有用。
此示例展示了能够保存任意类型值的单链表的简单类型声明。
作为练习，请为此链表的实现添加一些功能。

answer ref:
https://juejin.cn/post/7131387378429591582

方法中不能含有泛型，所以只能写为一般函数的形式。
我在这里实现了简单的增、删、查、打印，不过没有对可能的错误进行检查，你可以在这基础上继续完善。

这段代码实现了一个 泛型单链表，能够存储任意类型的数据，并提供了基本的增删查功能。
代码涉及到 Go 泛型、指针 和 链表结构，适合后端开发场景，比如缓存、任务队列等。

*/

// List 表示一个可以保存任何类型的值的单链表。List[T any]：定义了一个泛型的 单链表。
// T any 表示 T 可以是任意类型（int、string、struct 等）。  next *List[T]：指向下一个节点，实现链表结构。	val T：存储当前节点的值
type List[T any] struct {
	next *List[T]
	val  T
}

// 头部插入
func AddFront[T any](ls *List[T], v T) *List[T] {
	node := new(List[T])
	node.val = v
	node.next = ls.next
	ls.next = node
	return ls
}

// 尾部插入
func AddBack[T any](ls *List[T], v T) *List[T] {
	node := new(List[T])
	node.val = v
	ls.next = node
	return node
}

// 查找节点
func GetNode[T any](ls *List[T], v int) *List[T] {
	node := ls
	for i := 0; i <= v; i++ {
		node = node.next
	}
	return node
}

// 删除节点
func DeleteNode[T any](ls *List[T], v int) {
	node := GetNode(ls, v-1)
	node.next = node.next.next
}

// 遍历打印 Println
func Println[T any](ls *List[T]) {
	for node := ls.next; node != nil; node = node.next {
		print(node.val, " ")
	}
	print("\n")
}

func main() {
	head := new(List[int]) // 创建 int 类型的链表
	tail := head           // tail 指向 head
	tail = AddBack(tail, 1)
	tail = AddBack(tail, 5)
	tail = AddBack(tail, 7)
	head = AddFront(head, 3)
	head = AddFront(head, 6)
	Println(head)                 // 6 3 1 5 7
	println(GetNode(head, 2).val) // 1
	DeleteNode(head, 3)
	Println(head) // 6 3 1 7
}
