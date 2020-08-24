package tree

// 实现一个遍历函数
func (node *Node) Traverse() {
	// 使用中序遍历：中序遍历（LDR）是二叉树遍历的一种，也叫做中根遍历、中序周游。
	// 在二叉树中，中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树。
	// 判断空树的情况
	if node == nil {
		return
	}

	// 中序遍历
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}
