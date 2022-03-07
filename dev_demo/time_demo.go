package main

import (
	"fmt"
	"time"
)

// ref: https://www.jb51.net/article/164869.htm
// https://www.jb51.net/article/164869.htm
// [important] go常用时间/时间戳操作范例

func main() {

	date_time := time.Now()
	fmt.Print("01.获取当前时间: ")
	fmt.Println(date_time)
	fmt.Printf("获取当前时间类型：%T\n", date_time)
    fmt.Println("==================================")

	fmt.Println("02.获取 年 月 日 时 分 秒 纳秒")
	year := time.Now().Year() //  current year
	month := time.Now().Month()
	day := time.Now().Day()    // 日
	hour := time.Now().Hour()  // 小时
	minute := time.Now().Minute()  // 分钟
	second := time.Now().Second()  // 秒
	nanosecond := time.Now().Nanosecond() // 纳秒

	fmt.Println("年：", year)
	fmt.Println("月：", month)
	fmt.Println("日：", day)
	fmt.Println("小时：", hour)
	fmt.Println("分钟：", minute)
	fmt.Println("秒：", second)
	fmt.Println("纳秒：", nanosecond)

	fmt.Println("==================================")

	fmt.Println("03.获取当前时间戳")
	time_unix := time.Now().Unix()
	time_unix_milli := time.Now().UnixMilli()
	time_unix_micro := time.Now().UnixMicro()
	time_unix_nano := time.Now().UnixNano()

	fmt.Println("单位秒时间戳：", time_unix)  // 10位
	fmt.Println("单位毫秒时间戳：", time_unix_milli)  // 13位
	fmt.Println("单位微妙时间戳：", time_unix_micro)  // 16位
	fmt.Println("单位纳秒时间戳：", time_unix_nano)  // 19位

	fmt.Println("==================================")
	fmt.Println("04.当前时间戳格式化")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("==================================")

	fmt.Println("05.时间戳转为go格式的时间")
	var timeUnix int64 = 1563555859
	fmt.Println(time.Unix(timeUnix,0))
	// 之后可以用Format转换
	fmt.Println(time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05"))

	fmt.Println("==================================")

	fmt.Println("06.格式化字符串转时间戳")

	//从格式化字符串转为时间戳，第一个参数是格式，第二个参数是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/28/2022")
	fmt.Println("格式化字符串转为时间戳: ", tm2.Unix())

	//fmt.Println("==================================")
	//t := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local).Unix()
	//fmt.Printf("%T\n", t)
	//fmt.Println(t)

	fmt.Println("==================================")

	fmt.Println("07.时间戳转格式化字符串")
	timestamp := time.Now().Unix()
	// 格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println("时间戳转格式化字符串1: ", tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println("时间戳转格式化字符串2: ", tm.Format("02/01/2006 15:04:05 PM"))

}