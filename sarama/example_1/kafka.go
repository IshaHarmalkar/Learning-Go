package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func NewProducer(brokers []string) sarama.SyncProducer {

	config := sarama.NewConfig()
	config.Producer.Return.Successes  = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}

	return producer
}

func NewConsumer(brokers []string, groupID string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Failed to create kafka consumer group: %v", err)
	}

	return consumer

}

type SimpleConsumerHandler struct {}

func (SimpleConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil}

func (SimpleConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil}

func (h SimpleConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Received message: %s\n", string(msg.Value))
		sess.MarkMessage(msg, "")
	}

	return nil

}