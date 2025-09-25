package main

import (
	"Server/producer"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var userProducer *producer.Producer


func main(){

	//initialize the kafka producer
	var err error
	brokers := []string{"localhost:9092"}
	userProducer, err = producer.NewProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to initialize user producer: %v", err)
		os.Exit(1)
	}
		
	defer userProducer.Close()

	//http server routes
	http.HandleFunc("/users", userHandler)

	port := ":8080"

	log.Printf("Starting HTTP server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}


func userHandler(w http.ResponseWriter, r *http.Request){
	
	//only allow post
	if r.Method != http.MethodPost{
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	//json to user struct
	var user producer.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	partition, offset, err := userProducer.SendUserMessage("crud", user)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}


	//send success to client
	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User creation message sent to Kafka successfully!")




}