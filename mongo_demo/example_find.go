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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    _ = cancel
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
		fmt.Println(result)   //  [{_id ObjectID("635e396eaf684ef2a0e099d2")} {name pi} {value 3.141592653}]
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
