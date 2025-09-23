package main

import (
	"fmt"
	"kafkabasics/consumer"
	"kafkabasics/producer"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	bootstrapServers = "localhost:9092"
	topic            = "my-topic"
)

var p *kafka.Producer

func main(){
	var err error
	p, err = producer.InitProducer(bootstrapServers)

	if err != nil {
		log.Fatalf("COuld not initialize producer: %v", err)
	}

	defer p.Close()


	//start cinsuper in a separate foroutine
	go consumer.StartConsumer(bootstrapServers, topic)

	//set up http server
	http.HandleFunc("/produce", produceHandler)
	fmt.Println("Producer endpoint ready at http:/localhost:8080/produce?message=your_message")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


//http handler that sends a msg to kafka
	func produceHandler(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			http.Error(w, "message query parameter is required", http.StatusBadRequest)
			return
		}

		// Send the message to Kafka using the producer from the producer package
		err := producer.SendMessage(p, topic, message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			http.Error(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
			return
		}

		// Respond to the user
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Message '%s' sent successfully to Kafka!\n", message)
	}

func blockForever(){
	select {}
}