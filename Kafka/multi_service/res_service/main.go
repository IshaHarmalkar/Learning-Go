package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {

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
	partitionConsumer, err := consumer.ConsumePartition("test_event", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}

	//userProducer, err := NewProducer(brokers)

	defer partitionConsumer.Close()

	log.Println("Listening for messages..")


	db, err := NewUserRepository()
	if err != nil {
		log.Fatalf("Failed to iniatize user repository: %v", err)
	}

	for msg := range partitionConsumer.Messages(){

		var km KafkaMessage

		if err := json.Unmarshal(msg.Value, &km); err != nil{
			log.Panicf("Failed to deserialize user messaage: %v", err)
			continue
		}

		res, err := db.checkDuplicate(km)
		if err != nil {
			fmt.Println("checking for duplicates failed")
		}
		if res {
			fmt.Printf("Write to duplicates")
		} else {
			res, err := db.CreatePass(km)
			if err != nil {
				log.Panicf("Query failed: %v ", err)
				
			}

		  fmt.Println("written to db ", res)

		}

		
		



	}

}