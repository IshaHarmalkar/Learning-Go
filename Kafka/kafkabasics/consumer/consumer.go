package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker string, topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		GroupID: "my-consumer-group",
		Topic:   topic,
	})
	defer r.Close()

	fmt.Printf("Consumer started, waiting for messages on topic: %s\n", topic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(m.Value))
	}
}
