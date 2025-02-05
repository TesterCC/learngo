package main

import (
	"fmt"
	"math"
)

// P43

func main() {
	//高长度类型向低长度类型转换会发生长度截取，仅保留高长度类型的低位值，造成转换错误
	var e uint16 = math.MaxUint8 + 1
	fmt.Printf("e valud(unint16) is %v\n", e)
	var d = uint8(e)
	fmt.Printf("d value(unint8) is %v\n", d)

	// float type
	fmt.Printf("%f \n", math.E)   // 默认宽度和精度
	fmt.Printf("%.2f \n", math.E)  // 默认宽度和2位精度
}
