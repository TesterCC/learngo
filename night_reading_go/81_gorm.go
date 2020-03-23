package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
GORM 中文文档, Golang写的，开发人员友好的ORM库。
http://gorm.book.jasperxu.com/
https://www.bilibili.com/video/av97722329
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

	// 读取
	var product Product
	db.First(&product, 1)   // 查询id为1的product
	db.First(&product, "code = ?", "L1207")

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println(product)

	// 删除 - 删除product
	db.Delete(&product)
}
