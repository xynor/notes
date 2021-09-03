package types

import (
	"github.com/shopspring/decimal"
	"time"
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

func (d *Direction) String() string {
	switch *d {
	case Sell:
		return "Sell"
	case Buy:
		return "Buy"
	}
	return "UNKNOWN"
}

type Order struct {
	Uuid      string          `json:"uuid"`
	Oid       string          `json:"oid"`
	Action    Action          `json:"action"`
	Direction Direction       `json:"direction"`
	Symbol    string          `json:"symbol"`
	Qty       decimal.Decimal `json:"qty"`
	Price     decimal.Decimal `json:"price"`
	Time      time.Time       `json:"time"`
}

type Trade struct {
	Uuid     string          `json:"uuid"`
	Symbol   string          `json:"symbol"`
	Price    decimal.Decimal `json:"price"`
	Qty      decimal.Decimal `json:"qty"`
	TakerOid string          `json:"taker_oid"`
	MakerOid string          `json:"maker_oid"`
	Time     time.Time       `json:"time"`
}
