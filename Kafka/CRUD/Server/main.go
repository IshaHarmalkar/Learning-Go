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
	
	
	var  user producer.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var msgType string
	switch r.Method {
	case http.MethodPost:
		msgType = "create"
	case http.MethodPut:
		msgType = "update"
		if user.ID == 0 {
			http.Error(w, "Missing user ID for update", http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		msgType = "delete"
		if user.ID == 0 {
			http.Error(w, "Missing user ID for delete", http.StatusBadRequest)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	partition, offset, err := userProducer.SendCRUDMessage("crud", msgType, user)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s message sent to kafka successfully!", msgType)



}