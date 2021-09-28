package cmd

import (
	"github.com/spf13/cobra"
)

// 放置根命令 P23

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

// 每一个子命令，都是需要到 rootCmd 中进行注册的，否则将无法使用
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
