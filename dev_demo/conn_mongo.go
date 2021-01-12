package main

// ref: https://zhuanlan.zhihu.com/p/79585400
// ref: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
import (
	"context"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this User type later in the program
type Trainer struct {
	Name string
	Age int
	City string
}

/*
Run in Terminal:
go run conn_mongo.go
*/
func main() {
	// Rest of the code will go here

	// Step 1: 使用Go Driver连接到MongoDB
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 连接成功， 你可以通过将如下这行代码添加到main函数最后面来获得在crm库里面User集合的handle
	// CRUD by go mongo-driver
	collection := client.Database("test").Collection("trainers")

	lily := Trainer{"Lily", 30, "CS Town"}

	// 插入一个单独的文档， 使用 collection.InsertOne()函数
	insertResult, err := collection.InsertOne(context.TODO(), lily)  // _id自插入,所以可以多次插入

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert a single document: ", insertResult.InsertedID)

	// 不用时应该关闭连接
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
