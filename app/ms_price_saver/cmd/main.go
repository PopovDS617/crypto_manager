package main

import (
	"context"
	"log"
	"ms_price_saver/internal/app"
)

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	a.Run()

}
