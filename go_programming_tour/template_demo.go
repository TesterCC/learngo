package main

import (
	"os"
	"strings"
	"text/template"
)
/*
template是Go的文本模版引擎，提供2个标准库
- text/template: 基于模板输出文本内容
- html/template: 基于模板输出安全的HTML格式内容，进行了转义，以防注入攻击
 */
// template quick start
// 预定义的待解析模板
const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	// 通过 FuncMap 方法注册了名为title的自定义函数
	funcMap := template.FuncMap{"title":strings.Title}
	// 根据给定的名称标识，创建模板对象；调用Parse方法，将常量templateText(即与定义的待解析模板)解析为当前文本模版的主题内容
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}

	// 调用 Execute方法，进行模板渲染，即将传入的data动态参数渲染到对应的模板标识位上
	// 因为将 Execute的io.Writer指定到了os.Stdout中，所以最终输出到标准控制台
	_ = tpl.Execute(os.Stdout, data)
}
