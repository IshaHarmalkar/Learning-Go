package main

import (
	"log"

	"pokemon/internal/kafka"
)



func main() {

	brokers := []string{"localhost:9092"}
	groupID := "pokemon-consumer-group"
	topic := "pokemon_topic"

	if err := kafka.RunConsumer(brokers, groupID, topic); err != nil {
		log.Fatalf("consumer failed: %v", err)
	}

}