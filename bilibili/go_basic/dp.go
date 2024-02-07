package main

import (
	"fmt"
	"time"
)

// https://www.bilibili.com/video/BV11Q4y157um

// 2个长度为1000000的数组求求内积，向量内积计算, go version, python version in python script

func main() {
	const len2 int = 1000000 // 一百万
	var crr []float32
	var drr []float32

	for i := 0; i < len2; i++ {
		value := float32(i % 10)
		crr = append(crr, value)
		drr = append(drr, value)
	}

	begin := time.Now().UnixNano() // unit nanosecond

	prod := InnerProduct(crr, drr)

	end := time.Now().UnixNano()

	// convert nanosecond to second
	fmt.Printf("测试2个长度为1000000的数组求内积，向量内积计算）: %.0f\t\t用时%f秒\n", prod, float32(end-begin)/1e9)
}

func InnerProduct(x []float32, y []float32) float32 {
	// 2个长度为1000000的数组求求内积，向量内积计算
	var rect float32 = 0
	for i := 0; i < len(x); i++ {
		rect += x[i] * y[i]
	}
	return rect
}
