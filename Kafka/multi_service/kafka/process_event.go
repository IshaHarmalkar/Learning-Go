package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getUser(userId string) User {

	url := "http://localhost:8080/users/" + userId
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("https request failed:", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Converting to struct error: ", err)
	}

	return user
}

