package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* https://tour.go-zh.org/basics/1
每个 Go 程序都是由包构成的。

程序从 main 包开始运行。

本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
*/

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// 要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed (https://go-zh.org/pkg/math/rand/#Seed)
	rand.Seed(time.Now().Unix())
	fmt.Println("My Second favorite number is", rand.Intn(10))
	fmt.Println("My Second favorite number is", rand.Intn(3))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("My Second favorite number is", rand.Intn(10))
}
