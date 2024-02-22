package model

import "time"

type ReceivedTokenEntity struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Token struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type ProduceData struct {
	TokenData []Token
	Timestamp time.Time
}
