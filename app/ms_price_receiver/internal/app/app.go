package app

import (
	"encoding/json"
	"fmt"
	"ms_price_receiver/internal/client"
	"ms_price_receiver/internal/converter"
	"ms_price_receiver/internal/model"
	"ms_price_receiver/internal/producer"
	"time"
)

type App struct {
	httpClient   *client.HTTPClient
	dataProducer producer.DataProducer
}

func initHTTPClient(url string) *client.HTTPClient {
	return client.NewHTTPClient(url)
}

func initDataProducer(topic string) (producer.DataProducer, error) {
	return producer.NewDataProducer(topic)
}

func NewApp() (*App, error) {

	var err error

	url := `https://www.binance.com/api/v3/ticker/price?symbols=["BTCUSDT"]`
	topic := "prices"

	client := initHTTPClient(url)

	dataProducer, err := initDataProducer(topic)

	return &App{httpClient: client, dataProducer: dataProducer}, err
}

func (a *App) Run() {

	for {
		time.Sleep(time.Second * 2)

		res, err := a.httpClient.Client.Get(a.httpClient.Url, nil)

		responseTime := time.Now()

		if err != nil {
			fmt.Println(err)
		}

		var rawTokenData []model.ReceivedTokenEntity

		err = json.NewDecoder(res.Body).Decode(&rawTokenData)

		if err != nil {
			fmt.Println(err)
		}

		tokenData := converter.ToPriceFromReceivedTokenEntity(rawTokenData)

		produceData := model.ProduceData{
			TokenData: tokenData,
			Timestamp: responseTime}

		err = a.dataProducer.Produce(produceData)

		if err != nil {
			fmt.Println(err)
		}

	}

}
