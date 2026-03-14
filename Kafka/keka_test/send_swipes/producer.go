package main

import (
	"encoding/json"
	"fmt"
	"time"

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

func sendEvent(e Event)(error) {

	brokers := []string{"localhost:9092"}
	eventProducer, err := NewProducer(brokers)
	if err != nil {
		return fmt.Errorf("new producer could not be created in kafka log fn: %v", err)
	}

	defer eventProducer.Close()

	km := KafkaMessage{
		Event: e,	
	}	

	payload, err := json.Marshal(km)
	if err != nil {
		return fmt.Errorf("failed to marshal kafka msg to json inside sendEvent")
	}




	msg :=  &sarama.ProducerMessage{
		Topic: "test_swipes",
		Value: sarama.ByteEncoder(payload),
	}
 
	fmt.Println("waiting for 10 seconds and then sending swipe")
	time.Sleep(10 * time.Second)
   
   eventProducer.syncProducer.SendMessage(msg)
   fmt.Println("Swipe sent to topic")

   return nil

}