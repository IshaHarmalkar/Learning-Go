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


type User struct{
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type KafkaMessage struct {
	Action string `json:"type"` //create, update, del
	User User `json:"user"`
}

type Ack struct {
	Id int `json:"id"`
	Action string `json:"type"`

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



func (p *Producer) SendAck(ack Ack) (int32, int64, error){

	payload, err := json.Marshal(ack)

	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert to json when sending to consumer")
	}

	msg := &sarama.ProducerMessage{

		Topic : "acks",
		Value: sarama.ByteEncoder(payload),
	}

	fmt.Println("Ack sent")

	return p.syncProducer.SendMessage(msg)

}





