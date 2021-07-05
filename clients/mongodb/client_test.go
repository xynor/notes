package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Symbol string             `bson:"symbol"`
	Price  int64              `bson:"price"`
	Qty    int64              `bson:"qty"`
	Taker  int64              `bson:"taker"`
	Maker  int64              `bson:"maker"`
	Time   time.Time          `bson:"time"`
}

func Test_Basic(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://192.168.101.2:27017/order"))
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

	s, _ := time.Parse(time.RFC3339, "2021-07-01T12:52:00Z")
	e, _ := time.Parse(time.RFC3339, "2021-07-01T12:53:00Z")
	rs, err := collection.Aggregate(context.Background(), []bson.M{
		{"$match": bson.M{"time": bson.M{"$gte": s, "$lte": e}}},
		{"$sort": bson.M{"price": 1}},
		{"$group": bson.M{"_id": s, "min": bson.M{"$first": "$$ROOT"},
			"max": bson.M{"$last": "$$ROOT"}}},
	})
	if err != nil {
		t.Log("error:", err)
		return
	}
	if rs.Err() != nil {
		t.Log("error:", rs.Err())
		return
	}
	for rs.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		t.Log(rs.Current.String())
	}
	//Close the cursor once finished
	rs.Close(context.TODO())

	//order := Order{
	//	Symbol: "btc/usdt",
	//	Price:  1234,
	//	Qty:    1234,
	//	Taker:  1234,
	//	Maker:  32432,
	//	Time:   time.Now().UTC(),
	//}
	//insertResult, err := collection.InsertOne(context.TODO(), &order)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(insertResult)
}
