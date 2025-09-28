package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func startServer() {

	http.HandleFunc("/users", userHandler)

	port := ":8080"

	log.Printf("Starting http server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func userHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//db connection
	//databse credentials
	dsn := "root:@tcp(127.0.0.1:3306)/kafka_sync"

	//create the rpo instance for db operations
	userRepo, err := NewUserRepository(dsn)
	
	if err != nil {
		log.Fatalf("Failed to iniatize user repository: %v", err)
	}

	//var msgType string
	switch r.Method {

	case http.MethodPost:
		
		res, err := userRepo.CreateUser(u)
		handleError(err)		
		km, err := userRepo.LogKafkaMsg(res, "create")
		if err != nil {
			fmt.Println(err)
		}
		SendMsgToConsumer(km)
		ListenForAck()
		println("we are still in wswitch")

	case http.MethodPut:		
		if u.Id == 0 {
			http.Error(w, "Missing user ID for update", http.StatusBadRequest)
			return
		}

		res, err := userRepo.Update(u)
		handleError(err)
		km, err := userRepo.LogKafkaMsg(res, "update")
		if err != nil{
			fmt.Println(err)
		}

		SendMsgToConsumer(km)
		ListenForAck()

	case http.MethodDelete:		
		if u.Id == 0 {
			http.Error(w, "Missing user ID for delete", http.StatusBadRequest)
			return
		}
		res, err := userRepo.DeleteUser(u)
		handleError(err)
		km, err := userRepo.LogKafkaMsg(res, "delete")
		if err != nil{
			fmt.Println(err)
		}

		SendMsgToConsumer(km)
		ListenForAck()

		
		

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)


	

}


func handleError(err error){

	if err != nil {
		fmt.Println(err)

	}

}
