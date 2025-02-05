package sql2struct

import (
	"fmt"
	"learngo/go_programming_tour/cmd_tool/internal/word"
	"os"
	"text/template"
)

// 把得到的列信息按照特定的规则转换为Go结构体
// 采用模版渲染方案，写入预定义模版
const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

// 下面3个结构体，分别承担主轴的StructTemplate、StructColumn、StructTemplateDB
// StructColumn 用来存储转换后的Go结构体中的所有字段信息
// StructTemplateDB 用来存储最终用于渲染的模板对象信息
type StructTemplate struct {
	strcutTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

// 数据库类型到Go结构体的转换和对JSON Tag的处理
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}
// 对模块渲染的自定义函数和模块对象进行处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	// 声明了一个名为sql2struct的心模板对象
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
