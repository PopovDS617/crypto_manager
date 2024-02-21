package model

import "time"

type ReceivedTokenEntity struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Token struct {
	Symbol string
	Price  float64
}

type ProduceData struct {
	TokenData []Token
	Timestamp time.Time
}
