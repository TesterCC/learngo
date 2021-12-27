package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"   // 必须导入，然后进行MySQL驱动程序的初始化，否则会报错
)

// 1.4.3 表到结构体的转换 P48-56
// 想要获取标中列信息，需访问 information_schema 数据库中的 COLUMNS 表

// DBModel是整个数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 结构体 DBInfo 用于存储连接MySQL的一些基本信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// TableColumn 用于存储COLUMNS表中我们需要的一些字段
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 表字段类型映射
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

// 声明了初始化方法NewDBModel和三个核心结构体对象
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// P57 编写连接MySQL数据库的具体方法
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// debug
	fmt.Println(dsn)
	// 参数1 驱动名称 如mysql；参数2 驱动连接数据库的信息
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// 获取表中列的信息  针对COLUMNS表进行查询和数据组装
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
