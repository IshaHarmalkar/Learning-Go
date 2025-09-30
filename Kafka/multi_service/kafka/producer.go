package multiservice

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true


	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to start sarama producer: %w", err)
	}
	return &Producer{syncProducer: syncProducer}, nil
}

func (p *Producer) Close() error {
	return  p.syncProducer.Close()
} 

func sendEvent(km KafkaMessage)(error) {

	brokers := []string{"localhost:9092"}
	eventProducer, err := NewProducer(brokers)
	if err != nil {
		return fmt.Errorf("new producer could not be created in kafka log fn: %v", err)
	}

	defer eventProducer.Close()


	payload, err := json.Marshal(km)
	if err != nil {
		return fmt.Errorf("failed to marshal kafka msg to json inside sendEvent")
	}


	msg :=  &sarama.ProducerMessage{
		Topic: "test_event",
		Value: sarama.ByteEncoder(payload),
	}

   eventProducer.syncProducer.SendMessage(msg)

   return nil

}