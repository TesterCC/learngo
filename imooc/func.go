package main

import "fmt"

// 2-6 函数
/*
函数可返回多个值
函数返回多个值时可以起名字
仅用于非常简单的函数
对于调用者而言没有区别

多个返回值的使用场景通常是返回一个error
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


// 用函数式编程重写四则运算函数
// 2-6 10:19 #TODO

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(eval(3, 4, "x"))
	fmt.Println(div(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)
	a, b := div2(33, 3)
	fmt.Println(a, b)

	if result, err := eval(33,4,"x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
