package tree

import "fmt"

// https://www.bilibili.com/video/BV1u7411Y73E?p=16
// https://www.bilibili.com/video/BV1u7411Y73E?p=17

type Node struct {
	Value       int
	Left, Right *Node
}



// 工厂函数
// 1.使用自定义工厂函数   2.注意返回了局部变量的地址
func CreateNode(value int) *Node {
	return &Node{Value: value} // go语言局部变量也可以返回给别人用
}

// 为结构体定义方法，需要使用receiver接收者   这里是传值
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 注意，不是指针接收者的话，value还是为0，不是4
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}


