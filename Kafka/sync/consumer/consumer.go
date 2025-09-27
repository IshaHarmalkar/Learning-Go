package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {

	//kafka broker address
	brokers := []string{"localhost:9092"}

	//sarama configuration
	config := sarama.NewConfig()

	//create a consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start kafka consumer: %v", err)
	}

	defer consumer.Close()

	//consumer message from topic
	partitionConsumer, err := consumer.ConsumePartition("mangoes", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}

	defer partitionConsumer.Close()

	log.Println("Listening for messages..")

	//process messages as they get posted in the topic
	for message := range partitionConsumer.Messages() {
		log.Printf("Received message: %s", string(message.Value))
	}


}