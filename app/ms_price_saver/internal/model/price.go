package model

import "time"

type TokenPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type MessageQueueTokenDataList struct {
	TokenData []TokenPrice
	Timestamp time.Time
}

type MessageQueueTokenData struct {
	TokenData TokenPrice
	Timestamp time.Time
}
