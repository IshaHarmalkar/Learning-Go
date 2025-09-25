package producer

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type User struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type KafkaMessage struct{
	TYPE string `json:"type"` //create, update, delete
	User User   `json:"user"`
}

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


//send to kafka topic
func (p *Producer) SendCRUDMessage(topic string, msgType string, user User) (int32, int64, error){
	km := KafkaMessage{
		TYPE: msgType,
		User: user,
	}

	payload, err := json.Marshal(km)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to serialize message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
	}

	return p.syncProducer.SendMessage(msg)


}