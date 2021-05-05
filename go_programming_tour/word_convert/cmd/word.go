package cmd

import "github.com/spf13/cobra"

// 放置单词格式转换子命令word  P22

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {}

// P26 1.2.4 word 子命令
const (
	MODE_UPPER                         = iota + 1 // 全部单词转为大写
	MODE_LOWER                                    // 全部单词转为小写
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 下画线单词转为大写驼峰单词
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 下画线单词转为小写驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰单词转为下画线单词
)
