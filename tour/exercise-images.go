package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/*

http://127.0.0.1:3999/methods/25
https://tour.go-zh.org/methods/25

练习：图像
还记得之前编写的图片生成器 吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

ColorModel 应当返回 color.RGBAModel。

At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。

tour solution:

*/

type Image struct {
	Height, Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) At(x,y int) color.Color {
	c := uint8(x^y)
	return color.RGBA{c,c,255,255}
}

func main() {
	m := Image{256,256}
	pic.ShowImage(m)
}


