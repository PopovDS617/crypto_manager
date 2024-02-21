package main

import (
	"log"
	"ms_price_receiver/internal/app"
)

func main() {

	app, err := app.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	app.Run()

}
