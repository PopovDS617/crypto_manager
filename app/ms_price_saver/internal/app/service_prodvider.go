package app

import (
	"context"
	"log"
	"ms_price_saver/internal/client/db"
	"ms_price_saver/internal/client/db/pg"
	"ms_price_saver/internal/config"
	"ms_price_saver/internal/consumer"
	"ms_price_saver/internal/repository"
	tokenRepository "ms_price_saver/internal/repository/token"
)

type serviceProvider struct {
	pgConfig        config.PGConfig
	kafkaConfig     config.KafkaConfig
	dbClient        db.Client
	tokenRepository repository.TokenRepository
	dataConsumer    consumer.DataConsumer
}

func newServiceProvider() *serviceProvider {

	return &serviceProvider{}
}

func (s *serviceProvider) setPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}
func (s *serviceProvider) setKafkaConfig() config.KafkaConfig {
	if s.kafkaConfig == nil {
		cfg, err := config.NewKafkaConfig()
		if err != nil {
			log.Fatalf("failed to get kafka config: %s", err.Error())
		}

		s.kafkaConfig = cfg
	}

	return s.kafkaConfig
}

func (s *serviceProvider) setDBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.setPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) setTokenRepository(ctx context.Context) repository.TokenRepository {
	if s.tokenRepository == nil {
		s.tokenRepository = tokenRepository.NewRepository(s.setDBClient(ctx))
	}
	return s.tokenRepository
}

func (s *serviceProvider) setDataConsumer(_ context.Context) consumer.DataConsumer {
	if s.dataConsumer == nil {
		topic := s.setKafkaConfig().Topic()

		consumer, err := consumer.NewDataConsumer(topic)

		if err != nil {
			if err != nil {
				log.Fatalf("failed to create data consumer: %v", err)
			}
		}
		s.dataConsumer = consumer
	}

	return s.dataConsumer
}
