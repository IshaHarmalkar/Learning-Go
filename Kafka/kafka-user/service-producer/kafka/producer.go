package kafka

import (
	"context"
	"encoding/json"
	"service-producer/model"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}


func NewProducer(brokerAddress, topic string) *Producer{
	writer := &kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
	

	return &Producer{Writer: writer}

}


//sends a User message to Kafka
func(p *Producer) SendMessage(user model.User)error{
	data, err := json.Marshal(user)
	if err != nil{
		return err
	}

	msg := kafka.Message{
		Key: []byte(user.Name),
		Value: data,
	}

	return p.Writer.WriteMessages(context.Background(), msg)

}


//closes the kafka writer
func(p *Producer) Close() error {
	return p.Writer.Close()
}




NewProducer creates a Kafka writer for your topic (e.g., "user-created").

SendMessage marshals a User struct to JSON and sends it as a Kafka message.

Close should be called when your service stops to release resources.