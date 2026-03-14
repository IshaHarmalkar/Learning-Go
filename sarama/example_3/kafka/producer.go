package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return  &Producer{syncProducer: producer}, nil
}

func (p *Producer) SendMessage(topic, msg string) error {
	message := &sarama.ProducerMessage{
		Topic : topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.syncProducer.SendMessage(message)
	if err != nil {
		return err
	}

	fmt.Printf("[Producer] Message sent to topic %s, partition %d, offset %d\n", topic, partition, offset)
	return nil

}

func (p *Producer) Close() {
	if err := p.syncProducer.Close(); err != nil {
		log.Printf("Failed to close producer: %v", err)
	}
}