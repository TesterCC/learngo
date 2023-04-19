package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
一步步提升Go语言生成随机字符串的效率
https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go
*/

// 常见方法：二十六字母（大小写），然后随机取数，获得随机字符串。
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
    // 12 means str length
	randStr := RandStringRunes(12)
	fmt.Println(randStr)
}
