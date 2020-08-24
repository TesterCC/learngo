package main

import (
	"database/sql"
	"fmt"
)

// 1.4.3 表到结构体的转换 P56

// 结构体 DBInfo 用于存储连接MySQL的一些基本信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// DBModel是整个数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
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
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}
