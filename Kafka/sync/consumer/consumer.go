package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func main() {

	//kafka broker address
	brokers := []string{"localhost:9092"}

	//sarama configuration
	config := sarama.NewConfig()
	

	//create a producer
	ackProducer, err := NewProducer(brokers)

	if err != nil {
		log.Panicf("ack producer creation error")
	}

	//create a consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	
	if err != nil {
		log.Fatalf("Failed to start kafka consumer: %v", err)
	}


	defer consumer.Close()

	//consumer message from topic
	partitionConsumer, err := consumer.ConsumePartition("sync_user_test", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
  

	//userProducer, err := NewProducer(brokers)
	

	defer partitionConsumer.Close()

	log.Println("Listening for messages..")

	//process messages as they get posted in the topic
	for msg := range partitionConsumer.Messages() {
		log.Println("raw kafka msg: ", msg.Value)
		log.Println("Received message: ", string(msg.Value))
		
        var km KafkaMessage

		if err := json.Unmarshal(msg.Value, &km); err != nil{
			log.Panicf("Failed to deserialize user messaage: %v", err)
			continue
		}

		userId := km.User.Id
		action := km.Action

		ack := Ack{
			Id: userId,
			Action: action,
		}

		ackProducer.SendAck(ack)




		
	}


	



}

