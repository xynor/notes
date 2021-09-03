package engine

import (
	"container/list"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/log"
	"github.com/xinxuwang/ExchangeX/matchEngine/types"
	"testing"
)

func TestEngine_MatchBuyComplete(t *testing.T) {
	_ = log.Init(logrus.DebugLevel, "./matchEngine", "matchEngine")
	match := Engine{}
	match.BuyDepth.List = list.New()
	match.SellDepth.List = list.New()
	//Place Sell Order
	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Sell,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(1234),
		Price:     decimal.NewFromInt(100),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Sell,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(12),
		Price:     decimal.NewFromInt(50),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Buy,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(1234),
		Price:     decimal.NewFromInt(100),
	})
}

func TestEngine_MatchSellComplete(t *testing.T) {
	_ = log.Init(logrus.DebugLevel, "./matchEngine", "matchEngine")
	match := Engine{}
	match.BuyDepth.List = list.New()
	match.SellDepth.List = list.New()
	//Place Buy Order
	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Buy,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(12),
		Price:     decimal.NewFromInt(50),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Buy,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(1234),
		Price:     decimal.NewFromInt(100),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Sell,
		Symbol:    "btcusdt",
		Qty:       decimal.NewFromInt(1234),
		Price:     decimal.NewFromInt(100),
	})
}
