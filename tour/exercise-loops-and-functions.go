package main

import (
	"fmt"
	"math"
)

// http://127.0.0.1:3999/flowcontrol/8  https://tour.go-zh.org/flowcontrol/8
/*

https://tour.go-zh.org/flowcontrol/8
https://go.dev/tour/flowcontrol/8

ref: https://blog.csdn.net/qq_27818541/article/details/54345881

实现一个平方根函数：用牛顿法实现平方根函数。

计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：

z -= (z*z - x) / (2*z)
重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。

在提供的 func Sqrt 中实现它。无论输入是什么，对 z 的一个恰当的猜测为 1。 要开始，请重复计算 10 次并随之打印每次的 z 值。观察对于不同的值 x（1、2、3 ...）， 你得到的答案是如何逼近结果的，猜测提升的速度有多快。

提示：用类型转换或浮点数语法来声明并初始化一个浮点数值：

z := 1.0
z := float64(1)
然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数大于还是小于 10。 尝试改变 z 的初始猜测，如 x 或 x/2。你的函数结果与标准库中的 math.Sqrt 接近吗？

（*注：* 如果你对该算法的细节感兴趣，上面的 z² − x 是 z² 到它所要到达的值（即 x）的距离， 除以的 2z 为 z² 的导数，我们通过 z² 的变化速度来改变 z 的调整量。 这种通用方法叫做牛顿法。 它对很多函数，特别是平方根而言非常有效。）

*/

func Sqrt(x float64) float64 {
	z := 1.0    // 定义一个初始值并对它初始化
	temp := 0.0 // 临时变量，作为记录z上次的值

	for {
		z = z - (z*z-x)/(2*z) // 计算出最新的z值
		fmt.Println(z)
		if math.Abs(z-temp) < 0.000000000000001 {
			break //  当值停止改变（或改变非常小）的时候退出循环
		} else {
			temp = z //　赋值最终的结果
		}

	}
	return z
}

func main() {
	fmt.Println("牛顿法：", Sqrt(2))
	fmt.Println("math.Sqrt(２):", math.Sqrt(2))
}
