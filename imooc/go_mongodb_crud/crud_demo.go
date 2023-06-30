package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// 3-2 todo https://www.imooc.com/video/24258

func main() {
	// MongoDB connection config
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:toor@192.168.200.10:27017/admin"))

	if err != nil {
		fmt.Errorf("client connection estanblished failed, err: %v", err)
	}

	// timeout control
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create conn
	if err = client.Connect(ctx); err == nil {
		fmt.Println("connect to db success")
	}

	// conn db
	db := client.Database("test")

	// get collections
	colNames, err := db.ListCollectionNames(ctx, bson.M{})

	// output collection names
	fmt.Println("collection names: ", colNames)

	//// insert one data
	//insertDoc(client, ctx)

	// insert custom one data
	insertData := bson.M{"name": "Christine", "gender": "female"}
	insertOneDoc(client, ctx, insertData)

	// disconnect conn
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// write single line doc
// client is mongodb config , ctx is timeout value
func insertDoc(client *mongo.Client, ctx context.Context) {
	// connection db and col
	col := client.Database("test").Collection("insertonetest")

	// insert one data
	insertOneRet, err := col.InsertOne(ctx, bson.M{"name": "John", "gender": "male"})

	// erorr check
	if err != nil {
		log.Fatal(err)
	}

	// output insert_data _id
	fmt.Println("id: ", insertOneRet.InsertedID)
}

// extend, can use in work
func insertOneDoc(client *mongo.Client, ctx context.Context, bsonData bson.M) {
	// connection db and col
	col := client.Database("test").Collection("insertonetest")

	// insert one data
	insertOneRet, err := col.InsertOne(ctx, bsonData)

	// erorr check
	if err != nil {
		log.Fatal(err)
	}

	// output insert_data _id
	fmt.Println("id: ", insertOneRet.InsertedID)
}


// insert many data
func insertMany(client *mongo.Client, ctx context.Context)  {
	//  connect

}