package cmd

import (
	"github.com/spf13/cobra" // Cobra既是用于创建强大的现代CLI应用程序的库，也是用于生成应用程序和命令文件的程序。
	"learngo/go_programming_tour/word_convert/internal/word"
	"log"
	"strings"
)

var str string
var mode int8

// 放置单词格式转换子命令word  P22

var desc = strings.Join([]string{
	"该命令行工具支持各种单词格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下画线单词转为大写驼峰单词",
	"4：下画线单词转为小写驼峰单词",
	"5：驼峰单词转为下画线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("Output result: %s", content)
	},
}

// 单词内容和转换的模式进行命令行参数的设置和初始化
func init() {
	// 在 VarP 系列的方法中，第一个参数为需绑定的变量、第二个参数为接收该参数的完整的命令标志，第三个参数为对应的短标识，第四个参数为默认值，第五个参数为使用说明。
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

// P26 1.2.4 word 子命令
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)
