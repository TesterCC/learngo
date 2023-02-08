package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*
Go语言学习笔记—golang操作MongoDB数据库
ref:
https://blog.csdn.net/qq_39280718/article/details/126199197

use golang_db;
db.createCollection("student");

MongoDB中的JSON文档存储在名为BSON（二进制编码的JSON）的二进制表中。与其他将JSON数据存储为简单字符串和数字的数据库不同，BSON编码扩展了JSON表示，使其包含额外的类型，如int、long、date、浮点数和decimal128。这使得应用程序更容易可靠地处理、排序和比较数据。

连接MongoDB地Go驱动程序中有两大类型表示BSON数据：D和Raw。

类型D家族被用来简洁地构建使用本地Go类型地BSON对象。这对于构造传递给MongoDB地命令特别有用。D家族包括四大类：

D：一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongoDB命令。
M：一张无序的map。它和D是一样的，只是它不保持顺序。
A：一个BSON数组。
E：D里面的一个元素。

*/



var client *mongo.Client

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	var err error
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB 连接成功")
	client = c
}

func main() {
	initDB()
}
