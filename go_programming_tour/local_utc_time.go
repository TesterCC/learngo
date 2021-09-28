package main

import (
	"log"
	"time"
)

// 1.3.5.3 要注意的 time.Parse/Format
// https://golang2.eddycjy.com/posts/ch1/03-time2format/

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	inputTime := "2029-09-04 12:02:33"
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, inputTime)
	dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	log.Printf("输入时间：%s，输出时间：%s", inputTime, dateTime)
}
