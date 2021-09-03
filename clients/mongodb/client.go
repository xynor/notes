package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client struct {
	*mongo.Client
}

func NewMongoDB(uri string) (*Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: client,
	}, nil
}

//Types

type Status uint8

const (
	NotTrade Status = iota
	PartialTrade
	CompleteTrade
)

type Order struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Uuid   string               `bson:"uuid"`
	Oid    string               `bson:"oid"`
	Symbol string               `bson:"symbol"`
	Qty    primitive.Decimal128 `bson:"qty"`
	Price  primitive.Decimal128 `bson:"price"`
	Status Status               `bson:"status"`
	Time   time.Time            `bson:"time"`
}

type Trade struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Uuid     string               `bson:"uuid"`
	Symbol   string               `bson:"symbol"`
	Qty      primitive.Decimal128 `bson:"qty"`
	Price    primitive.Decimal128 `bson:"price"`
	TakerOid string               `bson:"taker_oid"`
	MakerOid string               `bson:"maker_oid"`
	Time     time.Time            `bson:"time"`
}
