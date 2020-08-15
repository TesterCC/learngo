package main

import (
	"os"
	"strings"
	"text/template"
)

// P53-P55
// 2种模板渲染方法
const templateText = `
Output 0: {{ title .Name1 }}
Output 1: {{ title .Name2 }}
Output 2: {{ .Name3 | title }}   
`

func main() {
	// 通过FuncMap方法注册了名为title的自定义函数
	funcMap := template.FuncMap{"title": strings.Title} // strings.Title 返回s中每个单词的首字母都改为标题格式的字符串拷贝
	//funcMap := template.FuncMap{"title": strings.ToLower} // strings.Title 返回s中每个单词的首字母都改为标题格式的字符串拷贝
	//funcMap := template.FuncMap{"title": strings.ToUpper} // strings.Title 返回s中每个单词的首字母都改为标题格式的字符串拷贝
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	//fmt.Println(tpl)
	//fmt.Println(a)  // nil

	data := map[string]string{
		"Name1": "GO",
		"Name2": "programming",
		"Name3": "tour",
	}

	_ = tpl.Execute(os.Stdout, data)
}
