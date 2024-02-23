package config

type PGConfig interface {
	DSN() string
}

type KafkaConfig interface {
	Topic() string
}
