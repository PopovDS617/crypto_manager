package consumer

import (
	"encoding/json"
	"fmt"
	"ms_price_saver/internal/model"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type DataConsumer interface {
	Consume() (*model.MessageQueueTokenDataList, error)
}

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewDataConsumer(topic string) (*KafkaConsumer, error) {

	bootstrapServers := os.Getenv("BOOTSTRAP_SERVERS")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Println("cannot subscribe to topic", err)
	}

	return &KafkaConsumer{
		consumer: c,
	}, nil
}

func (c *KafkaConsumer) Consume() (*model.MessageQueueTokenDataList, error) {
	var data model.MessageQueueTokenDataList
	msg, err := c.consumer.ReadMessage(-1)

	fmt.Println(msg)

	if err != nil {
		fmt.Printf("kafka consume error %s\n", err)
		return nil, err

	} else {

		if err := json.Unmarshal(msg.Value, &data); err != nil {
			fmt.Println(err)
			return nil, err
		}

	}
	return &data, nil
}
