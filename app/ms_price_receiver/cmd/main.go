package main

import (
	"fmt"
	"io"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func main() {

	url := `https://www.binance.com/api/v3/ticker/price?symbols=["BTCUSDT"]`

	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	for {
		time.Sleep(time.Second * 2)

		res, err := client.Get(url, nil)
		if err != nil {
			fmt.Println(err)
		}

		// Heimdall returns the standard *http.Response object
		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	}

}
