package main

import (
	"example_3/kafka"
	"fmt"
	"time"
)

func main() {

	topicA := "topic-a"
	topicB := "topic-b"

	producer, _ := kafka.NewProducer(kafka.Brokers)
	defer producer.Close()


	consumer, _ := kafka.NewConsumer(kafka.Brokers, topicB)
	go consumer.Start(func(msg string) {
		fmt.Printf("[Service B] Received: %s\n", msg)
	})

	counter := 1
	for {
		msg := fmt.Sprintf("Hello from Service B: %d", counter)
		if err := producer.SendMessage(topicA, msg); err != nil {
			fmt.Printf("[Service B] Failed to send: %v\n", err)
		}
		counter++
		time.Sleep(1 * time.Second)
	}



}