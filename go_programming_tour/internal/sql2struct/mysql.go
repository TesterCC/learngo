package main

import "database/sql"

// 1.4.3 表到结构体的转换 P56

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

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

