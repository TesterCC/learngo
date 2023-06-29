package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 3-2 todo https://www.imooc.com/video/24258

func main() {
	// MongoDB connection config
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:xxxx@localhost:27017/admin"))

	if err != nil {
		fmt.Errorf("client connection estanblished failed, err: %v", err)
	}

	// timeout control
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create conn
	if err = client.Connect(ctx); err == nil{
		fmt.Println("connect to db success")
	}

	// conn db
	db := client.Database("test")

	// get collections
	colNames, err := db.ListCollectionNames(ctx, bson.M{})

	// output collection names
	fmt.Println("collection names: ", colNames)

	// disconnect conn
}
