package producer

import (
	"encoding/json"
	"fmt"
	"ms_price_receiver/internal/model"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type DataProducer interface {
	Produce(model.ProduceData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewDataProducer(topic string) (*KafkaProducer, error) {
	bootstrapServers := os.Getenv("BOOTSTRAP_SERVERS")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (p *KafkaProducer) Produce(data model.ProduceData) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic,
			Partition: kafka.PartitionAny},
		Value: b,
	}, nil)

}
