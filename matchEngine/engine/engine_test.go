package engine

import (
	"container/list"
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/log"
	"github.com/xinxuwang/ExchangeX/matchEngine/types"
	"math/big"
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
		Qty:       big.NewInt(1234),
		Price:     big.NewInt(100),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Sell,
		Symbol:    "btcusdt",
		Qty:       big.NewInt(12),
		Price:     big.NewInt(50),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Buy,
		Symbol:    "btcusdt",
		Qty:       big.NewInt(1234),
		Price:     big.NewInt(100),
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
		Qty:       big.NewInt(12),
		Price:     big.NewInt(50),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Buy,
		Symbol:    "btcusdt",
		Qty:       big.NewInt(1234),
		Price:     big.NewInt(100),
	})

	match.Match(&types.Order{
		Action:    types.Create,
		Direction: types.Sell,
		Symbol:    "btcusdt",
		Qty:       big.NewInt(1234),
		Price:     big.NewInt(100),
	})
}
