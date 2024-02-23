package config

import (
	"errors"
	"os"
)

const (
	kafkaTopicEnvName = "KAFKA_TOPIC"
)

type kafkaConfig struct {
	topic string
}

func NewKafkaConfig() (KafkaConfig, error) {
	topic := os.Getenv(kafkaTopicEnvName)
	if len(topic) == 0 {
		return nil, errors.New("kafka config not found")
	}

	return &kafkaConfig{
		topic,
	}, nil
}

func (cfg *kafkaConfig) Topic() string {
	return cfg.topic
}
