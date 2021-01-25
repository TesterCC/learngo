package main

// ref: https://zhuanlan.zhihu.com/p/79585400
// ref: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
import (
	"context"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this User type later in the program
type Trainer struct {
	Name string
	Age  int
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

	// 1. insert document
	//lily := Trainer{"Lily", 30, "CS Town"}

	//ash := Trainer{"Ash", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Pewter City"}

	// 插入一个单独的文档， 使用 collection.InsertOne()函数
	//insertResult, err := collection.InsertOne(context.TODO(), lily)  // _id自插入,所以可以多次插入
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Insert a single document: ", insertResult.InsertedID)

	// 要同时插入多个文档， collection.InsertMany() 函数会采用一个slice对象
	//trainers := []interface{}{ash, misty, brock}
	//
	//insertRet, err := collection.InsertMany(context.TODO(), trainers)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted multiple documents: ", insertRet.InsertedIDs)

	// 2. update document
	// collection.UpdateOne()函数允许你更新单一的文档， 它需要一个filter文档来匹配数据库里面的文档， 并且需要一个update文档来描述更新的操作。 可以用bson.D类型来构建这些文档。

	/*
	    D系列的类型使用原生的Go类型简单地构建BSON对象。这可以非常有用的来创建传递给MongoDB的命令。 D系列包含4种类型：
		- D：一个BSON文档。这个类型应该被用在顺序很重要的场景， 比如MongoDB命令。
		- M: 一个无需map。 它和D是一样的， 除了它不保留顺序。
		- A: 一个BSON数组。
		- E: 在D里面的一个单一的子项。
	*/

	// collection.UpdateOne()函数允许你更新单一的文档， 它需要一个filter文档来匹配数据库里面的文档， 并且需要一个update文档来描述更新的操作。
	// 你可以用bson.D类型来构建这些文档
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 50}, // 这段代码会匹配name是“Ash”的文档， 并且将Ash的age增加 50
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// 4. 查找文档

	// 不用时应该关闭连接
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
