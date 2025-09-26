package kafka

import (
	"encoding/json"
	"fmt"
	"pokemon/internal/model"
	"strconv"

	"github.com/IBM/sarama"
)




type Producer struct {
	syncProducer sarama.SyncProducer
	topic   string
}

func NewProducer(brokers []string, topic string) (*Producer, error){
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

	p, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("create producer: %w", err)
	}

	return &Producer{syncProducer: p, topic: topic},nil

}


func (p *Producer) Close() error {
	return p.syncProducer.Close()
}

func (p *Producer) SendPokemon(pk model.Pokemon)(int32, int64, error){
	bytes, _ := json.Marshal(pk)
	key := sarama.StringEncoder(strconv.Itoa(pk.ID))
	msg := &sarama.ProducerMessage {
		Topic : p.topic,
		Key: key,
		Value: sarama.ByteEncoder(bytes),
	}

	return p.syncProducer.SendMessage(msg)
}


/* Wraps sarama kafka producer logic. Creating a producer, sending messages and
closing the connection. Configures partitioning strategy, random or keyed */
