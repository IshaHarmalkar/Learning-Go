package producer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func SendMessage(broker string, topic string, message string) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	fmt.Printf("Sent message: %s\n", message)
	return nil
}
