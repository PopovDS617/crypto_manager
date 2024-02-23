package app

import (
	"context"
	"fmt"
	"ms_price_saver/internal/converter"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	a.serviceProvider.setDataConsumer(ctx)
	a.serviceProvider.setTokenRepository(ctx)
	a.serviceProvider.setDBClient(ctx)

	return nil
}

func (a *App) Run() {

	for {

		msg, err := a.serviceProvider.dataConsumer.Consume()

		if err != nil {
			fmt.Println(err)
		} else {

			tokenData := converter.ToRepoFromMessageQueue(msg)

			for _, v := range tokenData {

				id, err := a.serviceProvider.tokenRepository.Create(context.Background(), &v)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("saved in DB with id: ", id)
				}

			}

		}

	}

}
