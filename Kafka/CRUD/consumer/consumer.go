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
	ID int `json:"id"`
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
		log.Fatalf("Failed to iniatize user repository: %v", err)
	}

	//configure Sarama consumer
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0  //actual version is 4.1.0, this is the closest one sarama supports
	config.Consumer.Return.Errors = true

	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil{
		log.Fatalf("Error creating cluster admin: %v", err)
	}
	defer admin.Close()

	topicName := "crud"
	topicDetail := &sarama.TopicDetail{
			NumPartitions:  1,
			ReplicationFactor: 1,
		
		}

	err = admin.CreateTopic(topicName, topicDetail, false)
	if err != nil {
		//"topic already exists is not fatal"
		if err.(*sarama.TopicError).Err == sarama.ErrTopicAlreadyExists{
			fmt.Printf("Topic %s already exists\n", topicName)
		}else {
			log.Fatalf("Error creating topic: %v", err)
		}
	} else {
		fmt.Printf("Topic %s created successfully\n", topicName)
	}


	//create a new Sarama consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil{
		log.Fatalf("Failed to start Sarama consumer: %v", err)
	}

	defer consumer.Close()

	//subsrice to the usr creations topic

	//partition consumer profides Messages channel and Errors channel.
	partitionConsumer, err := consumer.ConsumePartition("crud", 0, sarama.OffsetOldest)
	if err != nil{
		log.Fatalf("Failed to consume partition: %v", err)
	}

	defer partitionConsumer.Close()

	//HAndle OS signal for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	//start consuming message
	doneCh := make(chan struct{})

	//this is a go routine -> thread
	go func(){
		for{

			select{
			//new msg from kafka
			//partitionConsumer.Messages() is a channel of messages.
			case msg := <-partitionConsumer.Messages():
				//deserialize the message value (json) into a user struct
				var user User
				if err := json.Unmarshal(msg.Value, &user); err != nil{
					log.Panicf("Failed to deserialize user messaage: %v", err)
					continue
				}
				fmt.Printf("User created successfully: Name=%s, Email=%s\n", user.Name, user.Email)

				if err := userRepo.CreateUser(user); err != nil {
					log.Printf("Failed to create user in databse : %v", err)
				}

			//error from kafka
			case err := <-partitionConsumer.Errors():
				log.Panicf("Error from consumer: %v", err)

			//shutdown signal -> ctr + C 
			case <-signals:
				fmt.Println("Shutting down consumer...")
				close(doneCh)
				return
			}

		}
	}()

	<-doneCh


}


/* are go routines mostly processing some chanels?
	can channels be thought of as a stream that's constantly flooded with data?
	Is this process correct?
	1.Something is constantly happening
	2.we catch the said thing in a channel to be processed.const
	3. For each thing instnce usually a go routine is published
	4. the go routine will do it's logic and return something [or not] and the return thing
	will fall back into some other channel again or be caught in aother chanenl.const
*/