package cmd

import "github.com/spf13/cobra"

var calculateTime string
var duration string

// 创建time子命令 再到项目的 cmd/root.go 文件中进行相应的注册
var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {},
}

// 1.3.3.1 time now子命令 https://golang2.eddycjy.com/posts/ch1/03-time2format/