package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(bootstrapServers string, topic string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":"my-consumer-group",
		"auto.offset.reset": "earliest",

	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to subscribe to topic: %s\n", err)
		return
	}

	fmt.Printf("COnsumer started, waiting for message on topic: %s\n", topic)


	//channel for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			c.Close()
			return
		default:
			msg, err := c.ReadMessage(-1) // Wait indefinitely for a message
			if err == nil {
				fmt.Printf("Received message from Kafka: %s\n", string(msg.Value))
			} else if !err.(kafka.Error).IsFatal() {
				fmt.Printf("Consumer error: %v (ignoring)\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Consumer fatal error: %v\n", err)
				return
			}
		}
	}
}
