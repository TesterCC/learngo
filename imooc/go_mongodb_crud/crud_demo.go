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

	//// insert custom one data
	//insertData := bson.M{"name": "Christine", "gender": "female"}
	//insertOneDoc(client, ctx, insertData)

	//// insert many
	//insertMany(client, ctx)

	//// update data
	//updateData(client, ctx)

	// delete test
	findData(client, ctx)
	//deleteData(client, ctx)
	deleteManyData(client, ctx)
	fmt.Println("after delete, query again...")
	findData(client, ctx)

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
	// connect db and col
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
func insertMany(client *mongo.Client, ctx context.Context) {
	// connect db and col
	col := client.Database("test").Collection("insertmanytest")

	// define multi-lines slice
	docs := []interface{}{
		bson.M{"name": "test00", "gender": "male"},
		bson.M{"name": "test01", "gender": "female"},
		bson.M{"name": "test02", "gender": "male"},
	}

	// 自定应是否顺序写入，默认是true，如果写入多行数据时有一行报错，就会中断写入；如果是false时，会跳过报错，后面的继续执行。
	insertManyOpts := options.InsertMany().SetOrdered(false)

	// insert many lines
	insertManyResult, err := col.InsertMany(ctx, docs, insertManyOpts)

	//  error judgement
	if err != nil {
		log.Fatal(err)
	}

	// output all wriet _id
	fmt.Println("ids: ", insertManyResult.InsertedIDs)

}

// modify data
func updateData(client *mongo.Client, ctx context.Context) {
	// connect db and col
	col := client.Database("test").Collection("insertmanytest")

	// 如果过滤的文档不存在，则插入新的文档
	updateOneOpts := options.Update().SetUpsert(true)

	// update condition
	//updateOneFilter := bson.M{"name": "test00"}
	updateOneFilter := bson.M{"name": "xxx"}

	// after update
	//updateOneSet := bson.M{"$set": bson.M{"name": "Alfred"}}
	updateOneSet := bson.M{"$set": bson.M{"name": "Bob"}}

	// update
	updateResult, err := col.UpdateOne(ctx, updateOneFilter, updateOneSet, updateOneOpts)

	// check error
	if err != nil {
		log.Fatal(err)
	}

	// output update detail info
	fmt.Printf("matched: %d ; modified: %d ; upserted: %d ; upsertedID: %v\n",
		updateResult.MatchedCount,
		updateResult.ModifiedCount,
		updateResult.UpsertedCount,
		updateResult.UpsertedID, // when just update, here is nil
	)

}

// query mongodb data

func findData(client *mongo.Client, ctx context.Context) {
	// connect db and col
	col := client.Database("test").Collection("insertmanytest")

	// set order field, name field, update ascending order
	findOpts := options.Find().SetSort(bson.D{{"name", 1}})

	// query condition
	findCursor, err := col.Find(ctx, bson.M{"gender": "male"}, findOpts)

	// define result var
	var results []bson.M

	// query
	if err = findCursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	// loop for results
	for _, result := range results {
		fmt.Println(result)
	}

}

// delete an item data
func deleteData(client *mongo.Client, ctx context.Context) {
	// connect db and col
	col := client.Database("test").Collection("insertmanytest")

	// delete
	deleteResult, err := col.DeleteOne(ctx, bson.D{{"name", "test02"}})

	// check error
	if err != nil {
		log.Fatal(err)
	}

	// output delete lines count
	fmt.Println("delete count: ", deleteResult.DeletedCount)
}

// delete many data
func deleteManyData(client *mongo.Client, ctx context.Context) {
	// connect db and col
	col := client.Database("test").Collection("insertmanytest")

	// delete
	deleteManyResult, err := col.DeleteMany(ctx, bson.D{{"name", "test02"}})

	// check error
	if err != nil {
		log.Fatal(err)
	}

	// output delete lines count
	fmt.Println("delete count: ", deleteManyResult.DeletedCount)
}
