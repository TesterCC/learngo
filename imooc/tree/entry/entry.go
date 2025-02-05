package main

import (
	"fmt"
	"learngo/imooc/tree"
)

type myTreeNode struct {
	node *tree.Node
}


// 后序遍历（LRD）是二叉树遍历的一种，也叫做后根遍历、后序周游，可记做左右根。
// 后序遍历有递归算法和非递归算法两种。
// 在二叉树中，先左后右再根，即首先遍历左子树，然后遍历右子树，最后访问根结点。

/*
     3
0         5
   2   4
 */

// 实现后序遍历  通过组合方式实现
func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	// slice
	//nodes := [] Node {
	//	{value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	//fmt.Println(root)
	//
	//root.print()
	//fmt.Println()
	//root.Right.Left.setValue(4)
	//root.Right.Left.print()
	//fmt.Println()
	//
	//root.print() // 值接收者
	//root.setValue(100)

	//pRoot := &root // 是个地址
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()

	//var pRoot2 * Node
	//pRoot2.setValue(200)
	//pRoot2 = &root
	//pRoot2.setValue(300)
	//pRoot2.print()

}