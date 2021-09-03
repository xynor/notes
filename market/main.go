package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/clients/mongodb"
	"github.com/xinxuwang/ExchangeX/log"
	"github.com/xinxuwang/ExchangeX/market/market"
	"github.com/xinxuwang/ExchangeX/market/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {
	//Use http server for current dev
	_ = log.Init(logrus.DebugLevel, "./market", "market")
	client, _ := mongodb.NewMongoDB("mongodb://127.0.0.1:27017/order")
	mk := market.Market{
		MongoDBClient: client,
	}

	r := gin.Default()
	//curl -X POST -d '{"action":0,"direction":0,"symbol":"btcusdt","qty":"1234","price":"100","uuid":"1232432423","time":"2021-07-01T12:53:00Z"}' http://127.0.0.1:7777/placeOrder
	r.POST("/placeOrder", func(c *gin.Context) {
		logrus.Infoln("Recieved Order")
		var order types.Order
		data, _ := c.GetRawData()
		_ = json.Unmarshal(data, &order)
		symbol := order.Symbol
		//send to mongodb
		collectionName := order.Direction
		collection := mk.MongoDBClient.Database(fmt.Sprintf("%s-order", symbol)).Collection(collectionName.String())
		q, _ := primitive.ParseDecimal128(order.Qty.String())
		p, _ := primitive.ParseDecimal128(order.Price.String())
		or := mongodb.Order{
			Uuid:   order.Uuid,
			Oid:    order.Oid,
			Symbol: order.Symbol,
			Qty:    q,
			Price:  p,
			Time:   time.Time{},
		}
		_, err := collection.InsertOne(context.Background(), &or)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Recieved:%v", order.Uuid),
		})
	})
	r.Run(":7777")
}
