package main

import (
	"golang.org/x/tour/pic"
)

/*

http://127.0.0.1:3999/moretypes/18
https://tour.go-zh.org/moretypes/18

练习：切片
实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。

（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）

外层切片的长度为dy；
内层切片的长度为dx；
内层切片中的每个元素值为 (x+y)/2，x*y...；
使用嵌套循环的方式计算颜色值。
*/



func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy) // 创建外层切片，返回一个长度为 dy 的切片，元素类型为 uint8 的切片。
	for x := 0; x < dy; x++ {
		result[x] = make([]uint8, dx) // 创建里层切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。
		for y := 0; y < dx; y++ {
			result[x][y] = uint8(x * y) // //更改此句，可得到不同的结果
		}
	}
	return result

}

func main() {

	pic.Show(Pic)   // base64加密后输出的结果，需要将IMAGE:换成data:image/png;base64,  然后到http://tool.chinaz.com/tools/imgtobase解密。 或者 echo -n "" | base64 -D > testpic.png
}
