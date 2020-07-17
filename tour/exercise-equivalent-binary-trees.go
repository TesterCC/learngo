package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*
http://127.0.0.1:3999/concurrency/7
https://tour.go-zh.org/concurrency/7

http://127.0.0.1:3999/concurrency/8
https://tour.go-zh.org/concurrency/8

Exercise: Equivalent Binary Trees
等价二叉查找树

There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.
不同二叉树的叶节点上可以保存相同的值序列。例如，以下两个二叉树都保存了序列 `1，1，2，3，5，8，13`。

在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。

我们将使用 Go 的并发和信道来编写一个简单的解法。

本例使用了 tree 包，它定义了类型：

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

练习：等价二叉查找树
1. 实现 Walk 函数。

2. 测试 Walk 函数。

函数 tree.New(k) 用于构造一个随机结构的已排序二叉查找树，它保存了值 k, 2k, 3k, ..., 10k。

创建一个新的信道 ch 并且对其进行步进：

go Walk(tree.New(1), ch)
然后从信道中读取并打印 10 个值。应当是数字 1, 2, 3, ..., 10。

3. 用 Walk 实现 Same 函数来检测 t1 和 t2 是否存储了相同的值。

4. 测试 Same 函数。

Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1), tree.New(2)) 应当返回 false。

ref: Tree docs
https://godoc.org/golang.org/x/tour/tree#Tree
https://www.cnblogs.com/ahui2017/p/6383616.html?utm_source=itdadao&utm_medium=referral

tour solution:
https://github.com/golang/tour/blob/master/solutions/binarytrees_quit.go


函数 walkImpl 用递归方法遍历二叉树
*/

func walkImpl(t *tree.Tree,ch,quit chan int){
	if t == nil{
		return
	}
	walkImpl(t.Left,ch,quit)
	select {
	case ch <- t.Value:
		// Value successfully sent.
	case <- quit:
		return
	}
	walkImpl(t.Right,ch,quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree,ch, quit chan int){
	walkImpl(t,ch,quit)  //在 walkImpl 的外面再包裹一层 Walk 是为了 close(ch)
	close(ch)  // close(ch) 是为了判断这棵树是否已经遍历完毕，以便结束循环
}

// Same determines whether the trees.
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1,w2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1,w1,quit)
	go Walk(t2,w2,quit)

	for {
		v1, ok1 := <- w1
		v2, ok2 := <- w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}

}

func main() {

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1),tree.New(1)){
		fmt.Println("PASSED")
	}else{
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1),tree.New(2)){
		fmt.Println("PASSED")
	}else {
		fmt.Println("FAILED")
	}

}
