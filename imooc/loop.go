package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
2-5 循环

for,if 后面的条件没有括号
if条件里也可以定义变量
没有while
switch不需要break，也可以直接switch多个条件
*/

// 功能：把整数转成二进制数
func convertToBin(n int) string {
	result := ""
	if n==0 {
		return "0"
	}

	for ; n>0 ; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 功能：把整数转成二进制数v2
func convertToBin_v2(v int) string {
	result := ""

	if v==0 {
		return "0"
	}

	// 省略初始条件，相当于while
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}

	return result
}


// 用另一种方法read文件并print, 逐行读取
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {   //nil是一个重要的预先定义好的标识符，是许多种类型的零值表示  https://www.jianshu.com/p/174aa63b2cc5
		panic(err)   // 终止程序，然后报错，后面专门讲
	}

	scanner := bufio.NewScanner(file)

	//go语言中没有while，这样没有初始和终止条件的for语句，相当于while
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

func forever() {

	// 注意，这是一个死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1011 -> 1101
		convertToBin(777777),
		convertToBin(0),  // 空字符串没有显示，因为没有对0的处理
		)

	printFile("abc.txt")

	fmt.Println(
		convertToBin_v2(5), // 101
		convertToBin_v2(13), // 1011 -> 1101
		convertToBin_v2(777777),
		convertToBin_v2(0),  // 空字符串没有显示，因为没有对0的处理
	)

	//forever()
	
}

/*

cd ~/imooc.com/learngo/imooc

go run loop.go

*/