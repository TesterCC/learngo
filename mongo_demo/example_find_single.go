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

	var result struct {
		Value float64
	}
	filter := bson.D{{"name", "pi"}}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
