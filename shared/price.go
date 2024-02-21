package shared

type ReceivedTokenEntity struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type TokenPrice struct {
	Symbol string
	Price  float64
}
