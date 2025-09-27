package main

import (
	"encoding/json"
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


func userHandler(w http.ResponseWriter, r *http.Request){
	
	
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//db connection
	//databse credentials
	dsn := "root:@tcp(127.0.0.1:3306)/kafka"

	//create the rpo instance for db operations
	userRepo, err := NewUserRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to iniatize user repository: %v", err)
	}

	
	//var msgType string
	switch r.Method {

	case http.MethodPost:	
		userRepo.CreateUser(u)

		

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}




}