package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
GORM 中文文档, Golang写的，开发人员友好的ORM库。
http://gorm.book.jasperxu.com/
*/

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Connection database failed!")
	}
	defer db.Close()


	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1207", Price: 1000})

}
