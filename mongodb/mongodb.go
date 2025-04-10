package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"golang.org/x/net/context"
)

type (
	Map = map[string]any
)

func Coding() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("local")
	res, err := db.Collection("test").InsertOne(context.Background(), Map{
		"msg": "test",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(res.InsertedID)

	// count, err := db.Collection("test").CountDocuments(context.Background(), Map{
	// 	"test": "new",
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// coll := db.Collection("test")

	// //修改
	// res2, err := db.Collection("test").UpdateMany(ctx, filter, update)
	// // res2.ModifiedCount

	// res1, err := coll.DeleteOne(ctx, Map{})

	// result := Map{}

	// err = db.Collection("test").FindOne(context.Background(), Map{
	// 	"msg": "test",
	// }).Decode(&result)

	// if err != nil && err != mongo.ErrNoDocuments {
	// 	panic(err)
	// }

	// fmt.Println("result", result)

	// 设置查询选项
	opts := options.Find()
	opts.SetSkip(int64(0))  // 设置跳过的文档数量
	opts.SetLimit(int64(3)) // 设置返回的文档数量

	cursor, err := db.Collection("test").Find(context.Background(), Map{}, opts)

	if err != nil {
		fmt.Println("limit", err)
	}
	for cursor.Next(context.Background()) {
		var result Map
		if err := cursor.Decode(&result); err != nil {
			fmt.Println("limit", err)
		}
		fmt.Println("result", result)
	}

}
