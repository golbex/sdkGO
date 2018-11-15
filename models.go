package sdkGO

import (
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type Product struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Uid       string        `json:"uid"`
	First     string        `json:"first"`
	Second    string        `json:"second"`
	Increment float64       `json:"increment"`
	Min       float64       `json:"min"`
	Max       float64       `json:"max"`
}

type Candle struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Product   string        `json:"product"`
	Open      float64       `json:"open"`
	High      float64       `json:"high"`
	Low       float64       `json:"low"`
	Close     float64       `json:"close"`
	Timestamp int           `json:"timestamp" bson:"ts"`
	Volume    float64       `json:"volume"`
}

type NewOrder struct {
	Product string
	Type    string
	Side    string
	Price   float64
	Size    float64
}

type Order struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Status    string   `json:"status"`
	Type      string     `json:"type"`
	Side      string     `json:"side"`
	Product   string        `json:"product"`
	Price     float64       `json:"price"`
	Size      float64       `json:"size"`
	Fill      float64       `json:"fill"`
	UserID    bson.ObjectId `json:"user_id"`
	CreatedAt int           `json:"created_at"`
	ModifyAt  int           `json:"modify_at"`
	ClosedAt  int           `json:"closed_at"`
}

type Ticker struct {
	Ask     float64 `json:"ask"`
	Bid     float64 `json:"bid"`
	Price   float64 `json:"price"`
	Size    float64 `json:"size"`
	TradeId string  `json:"trade_id"`
	Vol24   float64 `json:"vol24"`
}
type Vol24 struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`
	Last   float64 `json:"last"`
}


type Stat struct {
	Ticker   Ticker `json:"ticker"`
	Volume24 Vol24  `json:"volume24"`
	sync.Mutex
}

type wallet struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	CurrencyCode string      `json:"currency_code"`
	Name         string      `json:"name"`
	PublicKey    string      `json:"public_key"`
	Value        walletValue `json:"value"`
}

type walletValue struct {
	In float64 `json:"in"` // внутренний баланс кошелька
}

type Error struct {
	Code int
	Msg  string
}