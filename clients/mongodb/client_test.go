package mongodb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_Basic(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/order"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database("order")
	collection := database.Collection("order")

	//s, _ := time.Parse(time.RFC3339, "2021-07-01T12:52:00Z")
	//e, _ := time.Parse(time.RFC3339, "2021-07-01T12:53:00Z")
	//rs, err := collection.Aggregate(context.Background(), []bson.M{
	//	{"$match": bson.M{"time": bson.M{"$gte": s, "$lte": e}}},
	//	{"$sort": bson.M{"price": 1}},
	//	{"$group": bson.M{"_id": s, "min": bson.M{"$first": "$$ROOT"},
	//		"max": bson.M{"$last": "$$ROOT"}}},
	//})
	//if err != nil {
	//	t.Log("error:", err)
	//	return
	//}
	//if rs.Err() != nil {
	//	t.Log("error:", rs.Err())
	//	return
	//}
	//for rs.Next(context.TODO()) {
	//	//Create a value into which the single document can be decoded
	//	t.Log(rs.Current.String())
	//}
	////Close the cursor once finished
	//rs.Close(context.TODO())

	p, _ := primitive.ParseDecimal128("0.000000000000000001")
	q, _ := primitive.ParseDecimal128("0.000000000000000001")
	order := Order{
		Symbol: "btc/usdt",
		Price:  p,
		Qty:    q,
		Time:   time.Now().UTC(),
	}
	insertResult, err := collection.InsertOne(context.TODO(), &order)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(insertResult)
}
