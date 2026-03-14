package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)
	

func main() {

   // fmt.Println("go works")
    broker := "localhost:9092"
    topic := "test-topic"

    println("test")
    // --- PRODUCER ---
    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer p.Close()

    // Produce a test message
    msg := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte("Hello Kafka from Go!"),
    }

    fmt.Println("Producing message...")
    p.Produce(msg, nil)

    // Wait for message deliveries
    p.Flush(15 * 1000)

    // --- CONSUMER ---
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": broker,
        "group.id":          "test-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }
    defer c.Close()

    c.SubscribeTopics([]string{topic}, nil)

    fmt.Println("Consuming message...")
    msg, err = c.ReadMessage(10 * time.Second)
    if err != nil {
        fmt.Printf("Consumer error: %v (%v)\n", err, msg)
    } else {
        fmt.Printf("Message received: %s\n", string(msg.Value))
    }
}
