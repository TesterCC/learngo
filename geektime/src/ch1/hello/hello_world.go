package main //表明代码所在模块
import (
	"fmt"
	"os"
) //引入代码依赖

/*
看课件
应用程序入口
1.必须是main包
2.必须是main方法
3.文件名不一定是main.go

与其它语言的主要差异：
1.Go中的main函数不支持任何返回值
2.通过ox.Exit来返回状态
3.main函数不支持传入参数
4.在程序中直接通过os.Args获取命令行参数
*/

//功能实现
func main() {
	fmt.Println(os.Args)   // go run hello_world.go test,在Terminal传入test
	if len(os.Args) > 1 {
		fmt.Println("Hello world!", os.Args[1])
	} else {
		fmt.Println("There is no args.")
	}

	os.Exit(0)    // go run in terminal
}
