package main

import (
	"encoding/json"
	"wahyu/dto"

	"github.com/google/uuid"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaServer = "localhost:9092"
	KafkaTopic  = "oders-v1-topic"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := KafkaTopic
	order := dto.Order {
		ID: uuid.New().String(),
		ProductId: uuid.New().String(),
		UserId: uuid.New().String(),
		Amount: 456000,
	}

	value, err := json.Marshal(order)

	if err != nil {
		panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny},
			Value: value,
	}, nil)

	if err != nil {
		panic(err)
	}
}