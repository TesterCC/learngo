package main

import (
	"fmt"
	"image"
)

/*

http://127.0.0.1:3999/methods/24
https://tour.go-zh.org/methods/24

image 包定义了 Image 接口

注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。
color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。

*/

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

