package main

import (
	"fmt"
	"log"
	"net/http"

	"kafkabasics/consumer"
	"kafkabasics/producer"
)

const (
	broker = "localhost:9092"
	topic  = "my-topic"
)

func main() {
	// Start consumer in a goroutine
	go consumer.StartConsumer(broker, topic)

	// HTTP producer endpoint
	http.HandleFunc("/produce", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			http.Error(w, "message query parameter is required", http.StatusBadRequest)
			return
		}

		err := producer.SendMessage(broker, topic, message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			http.Error(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Message '%s' sent successfully to Kafka!\n", message)
	})

	fmt.Println("Producer endpoint ready at http://localhost:8080/produce?message=your_message")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
