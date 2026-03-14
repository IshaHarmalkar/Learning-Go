package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	topic    string
	consumer sarama.Consumer
}

func NewConsumer(brokers []string, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		topic: topic,
		consumer: consumer,
	}, nil
}


//start consuming in a goroutine
func (c *Consumer) Start(onMessage func(msg string)) {

	partitions, err := c.consumer.Partitions(c.topic)
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}
/* 
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM) */
 
	for _, partition := range partitions {
		pc, err := c.consumer.ConsumePartition(c.topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to start consumer for partition %d: %v", partition, err)
		}

		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				onMessage(string(msg.Value))
			}
		}(pc)

	}

		/* <-signals
		fmt.Println("[Consumer] Stopped") */
}


func (c *Consumer) Close() {
	if err := c.consumer.Close(); err != nil {
		log.Printf("Failed to close cosnuemr: %v", err)
	}
}