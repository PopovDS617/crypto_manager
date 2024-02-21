package converter

import (
	"fmt"
	"ms_price_receiver/internal/model"
	"strconv"
)

func ToPriceFromReceivedTokenEntity(received []model.ReceivedTokenEntity) []model.Token {

	var result []model.Token

	for _, v := range received {

		parsedPrice, err := strconv.ParseFloat(v.Price, 32)

		if err != nil {
			fmt.Printf("unable to parse string %s to float\n", v.Price)
		} else {
			token := model.Token{Symbol: v.Symbol, Price: parsedPrice}
			result = append(result, token)
		}

	}

	return result
}
