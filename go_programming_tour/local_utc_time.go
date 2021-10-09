package main

import (
	"log"
	"time"
)

// 1.3.5.3 要注意的 time.Parse/Format
// https://golang2.eddycjy.com/posts/ch1/03-time2format/

func main() {
	// 通过标准库 time 中的 LoadLocation 方法来根据名称获取特定时区的 Location 实例
	location, _ := time.LoadLocation("Asia/Shanghai")   // 东八区 UTC+08:00  // 1
	inputTime := "2022-09-30 12:02:33"  // UTC time
	layout := "2006-01-02 15:04:05"      // 在 Go 语言中强调必须显示参考时间的格式，因此每个布局字符串都是一个时间戳的表示，并非随便写的时间点

	//// 实际上这与 Parse 方法有直接关系，因为 Parse 方法会尝试在入参的参数中中分析并读取时区信息，但是如果入参的参数没有指定时区信息的话，那么就会默认使用 UTC 时间。
	//t, _ := time.Parse(layout, inputTime)

	// 因此在这种情况下我们要采用 ParseInLocation 方法，指定时区就可以解决这个问题; 即所有解析与格式化的操作都最好指定时区信息
	t, _ := time.ParseInLocation(layout, inputTime, location)   // 2
	// 当你所部署的环境并不存在所设置时区的时区数据库时，也会导致 fallback 到 UTC 时区，因此与对接的运维人员确保部署时区的各方面设置是非常重要的。

	dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout) // 3 所有解析与格式化的操作都最好指定时区信息，否则当你遇到时区问题的时候，并且已经上线，那么后期再进行数据清洗就比较麻烦了。
	log.Printf("输入时间：%s，输出时间：%s", inputTime, dateTime)
}
