package engine

import (
	"container/list"
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/algo"
	"github.com/xinxuwang/ExchangeX/matchEngine/types"
)

type Engine struct {
	BuyDepth  algo.SortedList //for test
	SellDepth algo.SortedList //for test
}

func (e *Engine) Match(order *types.Order) {
	defer func() {
		logrus.Debugf("-PrintDepthS")
		e.PrintDepth(types.Buy)
		e.PrintDepth(types.Sell)
		logrus.Debugf("-PrintDepthE")
	}()

	if order.Direction == types.Buy {
		var n *list.Element
		for el := e.SellDepth.Front(); el != nil; el = n {
			preOrder := el.Value.(*types.Order)
			if preOrder.Price.Cmp(order.Price) <= 0 {
				logrus.Debugf("buy[price:%v,qty:%v] match sell[price:%v,qty:%v]", order.Price, order.Qty, preOrder.Price, preOrder.Qty)
				if preOrder.Qty.Cmp(order.Qty) <= 0 {
					logrus.Debugf("sell[price:%v,qty:%v] is taken all", preOrder.Price, preOrder.Qty)
					n = el.Next()
					e.SellDepth.Remove(el)
					order.Qty = order.Qty.Sub(order.Qty)
					logrus.Debugf("remain buy[price:%v,qty:%v]", order.Price, order.Qty)
					if order.Qty.Sign() <= 0 {
						return
					}
				} else {
					logrus.Debugf("buy[price:%v,qty:%v] is taken all", order.Price, order.Qty)
					preOrder.Qty = preOrder.Qty.Sub(order.Qty)
					logrus.Debugf("remain sell[price:%v,qty:%v]", preOrder.Price, preOrder.Qty)
					return
				}
			} else {
				logrus.Debugf("no more suitable sell,buy[price:%v,qty:%v]", order.Price, order.Qty)
				e.BuyDepth.Insert(order, types.DescendOrder)
				return
			}
		}
		e.BuyDepth.Insert(order, types.DescendOrder)
	} else if order.Direction == types.Sell {
		var n *list.Element
		for el := e.BuyDepth.Front(); el != nil; el = n {
			preOrder := el.Value.(*types.Order)
			if preOrder.Price.Cmp(order.Price) >= 0 && order.Qty.Sign() > 0 {
				logrus.Debugf("sell[price:%v,qty:%v] match buy[price:%v,qty:%v]", order.Price, order.Qty, preOrder.Price, preOrder.Qty)
				if preOrder.Qty.Cmp(order.Qty) <= 0 {
					logrus.Debugf("buy[price:%v,qty:%v] is taken all", preOrder.Price, preOrder.Qty)
					n = el.Next()
					e.BuyDepth.Remove(el)
					order.Qty = order.Qty.Sub(preOrder.Qty)
					logrus.Debugf("remain sell[price:%v,qty:%v]", order.Price, order.Qty)
					if order.Qty.Sign() <= 0 {
						return
					}
				} else {
					logrus.Debugf("sell[price:%v,qty:%v] is taken all", order.Price, order.Qty)
					preOrder.Qty = preOrder.Qty.Sub(order.Qty)
					logrus.Debugf("remain buy[price:%v,qty:%v]", preOrder.Price, preOrder.Qty)
				}
			} else {
				logrus.Debugf("no more suitable buy,sell[price:%v,qty:%v]", order.Price, order.Qty)
				e.SellDepth.Insert(order, types.AscendOrder)
				return
			}
		}
		e.SellDepth.Insert(order, types.AscendOrder)
	}
}

func (e *Engine) PrintDepth(d types.Direction) {
	if logrus.GetLevel() != logrus.DebugLevel {
		return
	}

	if d == types.Buy {
		for el := e.BuyDepth.Front(); el != nil; el = el.Next() {
			order := el.Value.(*types.Order)
			logrus.Debugf("BuyDepth price %v,qty %v", order.Price, order.Qty)
		}
	}

	if d == types.Sell {
		for el := e.SellDepth.Front(); el != nil; el = el.Next() {
			order := el.Value.(*types.Order)
			logrus.Debugf("SellDepth price %v,qty %v", order.Price, order.Qty)
		}
	}
}
