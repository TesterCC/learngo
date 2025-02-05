package main

import "fmt"

// https://www.runoob.com/go/go-basic-syntax.html
func main() {
	// %d 表示整型数字， %s 表示字符串
	var order_num = 77
	var end_time = "2021-07-07"
	var param = "order_num=%d&end_time=%s"

	var cat_str = fmt.Sprintf(param, order_num, end_time)
	fmt.Println(cat_str)
	//fmt.Print(cat_str)

}
