package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Create a producer
    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "test-topic"
    msg := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
        Value:          []byte("Hello from Go!"),
    }

    err = p.Produce(msg, nil)
    if err != nil {
        panic(err)
    }

    fmt.Println("Message produced successfully!")
}
