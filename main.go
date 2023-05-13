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

// https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// M is an unordered representation of a BSON document. This type should be used when the order of the elements does not
	// matter. This type is handled as a regular map[string]interface{} when encoding and decoding. Elements will be
	// serialized in an undefined, random order. If the order of the elements matters, a D should be used instead.
	// 是 BSON 文档的无序表示。当元素的顺序不正确时，应使用此类型
	// 在编码和解码时，此类型作为常规 map[string]interface{} 处理。元素将是
	// 以未定义的随机顺序序列化。如果元素的顺序很重要，则应改用 D。
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(databases)
}
