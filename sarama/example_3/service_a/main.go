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


	consumer, _ := kafka.NewConsumer(kafka.Brokers, topicA)
	consumer.Start(func(msg string) {
		fmt.Printf("[Service A] Received: %s\n", msg)
	})

	counter := 1

	for {
		msg := fmt.Sprintf("Hello from service A: %d", counter)
		if err := producer.SendMessage(topicB, msg); err != nil {
			fmt.Printf("[Service A] Faied to send: %v\n", err)
		}
		
		counter++
		time.Sleep(1 * time.Second)
		

	}
}


