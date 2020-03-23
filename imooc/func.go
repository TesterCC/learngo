package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 2-6 函数
/*
函数可返回多个值
函数返回多个值时可以起名字
仅用于非常简单的函数
对于调用者而言没有区别

多个返回值的使用场景通常是返回一个error

返回值的类型写在最后面
函数可以作为参数
没有默认参数，可选参数
 */

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b) // _为占位符，因为go语言编译要求变量一定要用到
		return q, nil
	default:
		//return panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 实现带余数除法，e.g.: 13 / 3 = 4 ... 1
// q表示商，r表示余数
//func div(a,b int) (int,int) {
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r // 函数体长就不推荐这么使用了。
}

// 用函数式编程重写四则运算函数  定义复合函数
// op func(int,int) int -> 定义函数op，op接受2个int参数，函数op返回一个int
func apply(op func(int, int) int, a, b int) int {

	// 获取函数名称
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)", opName, a, b)

	return op(a, b)
}

// rewrite pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// go语言有可变参数列表  ...int表示随便传入多少个int都可以
/*
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(eval(3, 4, "x"))
	fmt.Println(div(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)
	a, b := div2(33, 3)
	fmt.Println(a, b)

	if result, err := eval(33, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// 此方法代码阅读性好
	fmt.Println(apply(pow, 3, 4))  // 3的4次方

	// 也可以直接定义一个匿名函数传入，但是代码阅读性差
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1,2,3,4,5,6,7))  // 28
}
