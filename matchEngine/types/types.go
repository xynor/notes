package types

import (
	"math/big"
)

type Action uint8

const (
	Create Action = iota
	Delete
)

type Direction uint8

const (
	Sell Direction = iota
	Buy
)

type Order struct {
	Action    Action    `json:"action"`
	Direction Direction `json:"direction"`
	Symbol    string    `json:"symbol"`
	Qty       *big.Int  `json:"qty"`
	Price     *big.Int  `json:"price"`
}

func AscendOrder(v, current interface{}) bool {
	return v.(*Order).Price.Cmp(current.(*Order).Price) <= 0 //ascend
}

func DescendOrder(v, current interface{}) bool {
	return v.(*Order).Price.Cmp(current.(*Order).Price) >= 0 //descend
}
