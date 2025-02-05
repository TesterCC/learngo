package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://github.com/mongodb/mongo-go-driver
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect
func main() {
	// Create a direct connection to a host. The driver will send all requests
	// to that host and will not automatically discover other hosts in the
	// deployment.

	clientOpts := options.Client().ApplyURI(
		"mongodb://localhost:27017/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	//_ = client
	collection := client.Database("scc").Collection("demo")

	now := time.Now()
	//_id := int32(now.Unix())  // 10-digit stamp
	_id2 := int64(now.UnixMilli())  // 13-digit stamp

	bsonDict := bson.D{
		{"_id", _id2},
		{"name", "def"},
		{"value", 456},
	}

	res, err := collection.InsertOne(context.Background(), bsonDict)
	if err != nil {
		fmt.Println(err)
	}
	id := res.InsertedID

	fmt.Println(id)

}
