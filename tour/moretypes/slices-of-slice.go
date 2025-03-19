package main

import (
	"fmt"
	"strings"
)

/*
https://tour.go-zh.org/moretypes/14
https://go.dev/tour/moretypes/14

切片中的切片

切片可以包含任何类型，当然也包括其他切片。
*/

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		// index 0 1 2
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board) // [[_ _ _] [_ _ _] [_ _ _]]

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
