package main

/*
 make kafka msg and sent to consumer
*/

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

//New producer initializes and returns a new producer instance
func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	

	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to start Sarama producer: %w", err)

	}
	return &Producer{syncProducer: syncProducer}, nil
}

//shuts down the producer
func(p *Producer) Close() error {
	return p.syncProducer.Close()
}



func SendMsgToConsumer(KafkaMessage KafkaMessage)(error) {


	 //var producerPtr *producer.Producer
	brokers := []string{"localhost:9092"}

	userProducer, err := NewProducer(brokers)

	if err != nil {
		return fmt.Errorf("new producer could not be created in kafka log fn: %v", err)
	}	

	defer userProducer.Close()



	payload, err := json.Marshal(KafkaMessage)

	if err != nil {
	  return fmt.Errorf("failed to convert to json when sending to consumer")
	}

	msg := &sarama.ProducerMessage{
		Topic : "sync_user_test",
		Value: sarama.ByteEncoder(payload),
	}

	userProducer.syncProducer.SendMessage(msg)

	return nil



}
