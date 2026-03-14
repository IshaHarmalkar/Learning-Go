package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"

	"sync/shared"
)

/*
 make kafka msg and sent to consumer
*/



func SendMsgToConsumer(KafkaMessage KafkaMessage)(error) {

	 //var producerPtr *producer.Producer
	brokers := []string{"localhost:9092"}
	userProducer, err := NewProducer(brokers)
	if err != nil {
		return fmt.Errorf("new producer could not be created in kafka log fn: %v", err)
	}	

	defer userProducer.Close()



	payload, err := json.Marshal(KafkaMessage)
	handleError(err)
	msg := &sarama.ProducerMessage{
		Topic : "sync_user_test",
		Value: sarama.ByteEncoder(payload),
	}
	userProducer.syncProducer.SendMessage(msg)
	return nil



}
