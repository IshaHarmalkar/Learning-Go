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


//serializes a user and sends it to a kafka topic

func(p *Producer) SendUserMessage(topic string,user User)(int32, int64, error) {
	//serialize the user struct to json to send to kafka

	userJSON, err := json.Marshal(user)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to serialize user to json: %w", err)
	}

	//create a kafka msg
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(userJSON),
	}

	//sends msg to kafka topic
	partition, offset, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to send message to Kafka: %w", err)
	}

	return partition, offset, nil


}



func(p *Producer) SendUpdateMessage(topic string, user User) (int32, int64, error){
	return p.SendUserMessage(topic, user)
}


func (p *Producer) SendDeleteMessage(topic string, userID int) (int32, int64, error){

	payload := map[string]int{"id": userID}
	userJSON, err := json.Marshal(payload)

	if err != nil {
		return 0, 0, fmt.Errorf("failed to seialize delete message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic : topic,
		Value: sarama.ByteEncoder(userJSON),
	}

	return p.syncProducer.SendMessage(msg)

}