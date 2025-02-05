package main

import "fmt"

// 2-4 条件语句
/*
 1.if的条件里不需要括号
 */

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func main() {

	fmt.Println(bounded(200))
	fmt.Println(bounded(-20))
	fmt.Println(bounded(77))

}
