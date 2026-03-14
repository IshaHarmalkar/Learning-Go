package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Create producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	fmt.Println("Producer started")

	// Delivery report handler
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

	// Produce messages
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "Confluent", "Kafka", "Go", "client"} {
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)

		if err != nil {
			fmt.Printf("Failed to produce message: %v\n", err)
		}
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
}
