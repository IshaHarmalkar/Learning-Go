package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)



type Ack struct {
	Id int `json:"id"`
	Action string `json:"type"`

}




func ListenForAck() {

	println("New consumer")

	//kafka broker address
	brokers := []string{"localhost:9092"}

	//sarama configuration
	config := sarama.NewConfig()

	//create a consumer
	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		log.Fatalf("Failed to start kafka consumer: %v", err)
	}

	defer consumer.Close()

	//consumer message from topic
	partitionConsumer, err := consumer.ConsumePartition("acks", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}

	defer partitionConsumer.Close()


	log.Println("Listening for ack messages..")

	//process messages as they get posted in the topic
	for msg := range partitionConsumer.Messages() {
		
		
		var ackMsg Ack

		if err := json.Unmarshal(msg.Value, &ackMsg); err != nil{
			log.Panicf("Failed to deserialize user messaage: %v", err)
			continue
		}

		userId := ackMsg.Id
		//action := ackMsg.Action

		dsn := "root:@tcp(127.0.0.1:3306)/kafka_sync"

		//create the rpo instance for db operations
		ackRepo, err := NewUserRepository(dsn)
		if err != nil {
			log.Fatalf("Failed to iniatize user repository: %v", err)
		}

		fmt.Println("ack sent for processing")
		ackRepo.synAck(userId)
		
	}
}