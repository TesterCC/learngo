package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// 要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed (https://go-zh.org/pkg/math/rand/#Seed)
	rand.Seed(time.Now().Unix())
	fmt.Println("My Second favorite number is", rand.Intn(10))
}
