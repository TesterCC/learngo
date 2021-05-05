package cmd

import "github.com/spf13/cobra"

// 放置根命令 P23

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
