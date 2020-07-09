package main

// https://www.bilibili.com/video/BV18Q4y1M7NV?p=15

type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
}
