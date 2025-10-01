package main

import (
	"io"
	"log"
	"net/http"
)

func getUser() {

	resp, err := http.Get("http://localhost:8080/users/1")
	if err != nil {
		log.Fatalln("https request failed:", err)
	}


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	//convert the body to str
	sb := string(body)
	log.Printf(sb)

}