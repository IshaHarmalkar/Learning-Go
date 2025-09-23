package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// InitProducer initializes and returns a new Kafka producer.
func InitProducer(bootstrapServers string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	fmt.Printf("Successfully created producer: %s\n", p.String())
	return p, nil
}

// sends a message to the specified Kafka topic.
func SendMessage(p *kafka.Producer, topic string, message string) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return fmt.Errorf("produce failed: %w", err)
	}

	// Wait for the message to be delivered or for an error to occur.
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
	}

	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}
