package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

func main() {

	brokers := []string{"localhost:9092"}
	groupID := "example-group"
	topic := "test_topic"

	//create producer
	producer := NewProducer(brokers)
	defer producer.Close()

	//send a msg
	msg := &sarama.ProducerMessage {
		Topic: topic,
		Value: sarama.StringEncoder("Hello from producer!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Printf("Message sent to partition %d at offset %d", partition, offset)

	consumer := NewConsumer(brokers, groupID)
	defer consumer.Close()

	//start consuming messages
	ctx := context.Background()
	topics := []string{topic}
	handler := SimpleConsumerHandler{}

	log.Println("Consumer started...")
	if err := consumer.Consume(ctx, topics, handler); err != nil {
		log.Fatalf("Error while consuming: %v", err)
	}




}