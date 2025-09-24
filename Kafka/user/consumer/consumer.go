package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	//kafka broker address
	brokers := []string{"localhost:9092"}

	//databse credentials
	dsn := "root:@tcp(127.0.0.1:3306)/kafka"

	//create the rpo instance for db operations
	userRepo, err := NewUserRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to iniatize user reository: %v", err)
	}

	//configure Sarama consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	//create a new Sarama consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil{
		log.Fatalf("Failed to start Sarama consumer: %v", err)
	}

	defer consumer.Close()

	//subsrice to the usr creations topic
	partitionConsumer, err := consumer.ConsumePartition("user_creations", 0, sarama.OffsetOldest)
	if err != nil{
		log.Fatalf("Failed to consume partition: %v", err)
	}

	defer partitionConsumer.Close()

	//HAndle OS signal for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	//start consuming message
	doneCh := make(chan struct{})
	go func(){
		for{

			select{
			case msg := <-partitionConsumer.Messages():
				//deserialize the message value (json) into a user struct
				var user User
				if err := json.Unmarshal(msg.Value, &user); err != nil{
					log.Panicf("Failed to deserialize user messaage: %v", err)
					continue
				}
				fmt.Printf("User created successfully: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)

				if err := userRepo.CreateUser(user); err != err {
					log.Printf("Failed to create user in databse : %v", err)
				}

			case err := <-partitionConsumer.Errors():
				log.Panicf("Error from consumer: %v", err)

			case <-signals:
				fmt.Println("Shutting down consumer...")
				close(doneCh)
				return
			}

		}
	}()

	<-doneCh


}