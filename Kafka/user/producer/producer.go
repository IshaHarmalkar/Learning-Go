package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	brokers := []string{"localhost:9092"}

	//create a new sarama producer

	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil{
		log.Fatalf("Failed to start Sarma producer: %v", err)
	}
	defer producer.Close()

	//create a new user
	user := User{
		ID: 1,
		Name: "JOhn Doe",
		Email: "john.doe@example.com",
	}

	//serialie the user to json
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Failed to serialize user to JSON: %v", err)
	}

	//Create a Kafka Message
	msg := &sarama.ProducerMessage{
		Topic: "user_creations",
		Value: sarama.ByteEncoder(userJSON),
	}

	//send message to kafka topic
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %f\n", partition, offset)

}