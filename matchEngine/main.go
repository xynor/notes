package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/log"
	"github.com/xinxuwang/ExchangeX/matchEngine/engine"
	"github.com/xinxuwang/ExchangeX/matchEngine/types"
)

func main() {
	//Use http server for current dev
	_ = log.Init(logrus.DebugLevel, "./matchEngine", "matchEngine")
	match := engine.Engine{}
	match.BuyDepth.List = list.New()
	match.SellDepth.List = list.New()

	r := gin.Default()
	//curl -X POST -d '{"action":0,"direction":0,"symbol":"btcusdt","qty":1234,"price":100}' http://127.0.0.1:9999/sendOrder
	r.POST("/sendOrder", func(c *gin.Context) {
		logrus.Infoln("Recieved Order")
		var order types.Order
		data, _ := c.GetRawData()
		_ = json.Unmarshal(data, &order)
		match.Match(&order)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Recieved:%v", order.Price),
		})
	})
	r.Run(":9999")
}
