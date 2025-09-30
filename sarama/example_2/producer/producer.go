package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {

	brokers := []string{"localhost:9092"}
	topic := "test_topic"

	//create a new new synchronous producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start sarama producer: %v", err)
	}

	defer producer.Close()

	//msg to send

	messages := []string{"Hello Kafka!", "This is a sarama test", "apple is red"}

	for _, msg := range messages {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(msg),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}


	}


}