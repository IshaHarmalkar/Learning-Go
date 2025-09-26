package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pokemon/internal/model"
	"syscall"

	"github.com/IBM/sarama"
)


func ensureTopic(brokers []string, topic string, partitions int32) error {
	cfg := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(brokers, cfg)

	if err != nil {
		return fmt.Errorf("create admin: %w", err)
	}

	defer admin.Close()

	topics, err := admin.ListTopics()
	if err != nil {
		return fmt.Errorf("list topics: %w", err)
	}

	if _, exists := topics[topic]; exists {
		log.Printf("topic %s already exists", topic)
		return nil
	}

	detail := &sarama.TopicDetail{
		NumPartitions:  partitions,
		ReplicationFactor: 1,
	}

	if err := admin.CreateTopic(topic, detail, false); err != nil {
		return fmt.Errorf("create topic: %w", err)
	}

	log.Printf("topic %s created with %d partitions", topic, partitions)
	return  nil
}


type handler struct{}

func(h *handler) Setup(_ sarama.ConsumerGroupSession) error {return nil }
func(h *handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil}

func (h *handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var pk model.Pokemon
		if err := json.Unmarshal(msg.Value, &pk); err != nil {
			log.Printf("bad messafe: %v", err)
			sess.MarkMessage(msg, "")
			continue
		}
		log.Printf("[consumer] partition=%d offset=%d key=%s -> %+v",
			msg.Partition, msg.Offset, string(msg.Key), pk)
		sess.MarkMessage(msg, "")

	}
	return  nil
}

func RunConsumer(brokers []string, groupID string, topic string) error {

	if err := ensureTopic(brokers, topic, 3); err != nil {
		return err
	}

	cfg := sarama.NewConfig()
	cfg.Version = sarama.V3_6_0_0 //closest version to kafka
	cfg.Consumer.Return.Errors = true

	client, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil{
		return fmt.Errorf("consumer group: %w", err)
	}

	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			if err := client.Consume(ctx, []string{topic}, &handler{}); err != nil {
				log.Printf("consume error: %v", err)
			}
			if ctx.Err() != nil {
				return 
			}

			
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("shutting down consumer")
	return nil




}