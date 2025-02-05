package cmd

// 需要将其集成到我们的子命令中，也就是创建项目的 time 子命令
import (
	"github.com/spf13/cobra"
	"learngo/go_programming_tour/cmd_tool/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

// 创建time子命令 再到项目的 cmd/root.go 文件中进行相应的注册
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

//完成 time 子命令编写后，再到项目的 cmd/root.go 文件中进行相应的注册即可
// 1.3.3.1 time now子命令 https://golang2.eddycjy.com/posts/ch1/03-time2format/

// 针对获取当前时间，我们在 time 子命令下再新增一个 now 子命令，用于处理其具体的逻辑
// 创建nowTime子命令
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

/*
在获取当前时间的 Time 对象后，一共输出了两个不同格式的时间，分别如下：

第一个格式：通过调用 Format 方法设定约定的 2006-01-02 15:04:05 格式来进行时间的标准格式化。

第二个格式：通过调用 Unix 方法返回 Unix 时间，就是我们通俗说的时间戳 timestamp，其值为自 UTC 1970 年 1 月 1 日起经过的秒数。

如果你想要定义其它时间格式的话，标准库 time 中还支持（内部预定义）如下格式:

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	...
)

可以像这样子使用这些预定义格式，例如：

t := time.Now().Format(time.RFC3339)

*/

// 针对时间推算的处理
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

// 支持了三种常用时间格式的处理，分别是：时间戳、2006-01-02 以及 2006-01-02 15:04:05。 // strings.Count() 为作者更新后的方法

// 对 time 子命令进行 now、calc 的子命令和所需的命令行参数进行注册即可
func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", ` 需要计算的时间，有效单位为时间戳或已格式化后的时间 `)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", ` 持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
}