package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "[^\u4e00-\u9fa5]不可能[^\u4e00-\u9fa5]不可能$"

	s := "不可能"
	r := regexp.MustCompile(s)

	fmt.Println(r.Match([]byte(pattern)))

}
