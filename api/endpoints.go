package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)





type Payload struct {
	Pagination Pagination `json:"pagination"`
}


type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}


func getSites(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("JWT_TOKEN")

	getSitesUrl := os.Getenv("POST_SITES")

	payload := Payload{
		Pagination: Pagination{
			Page:    1,
			PerPage: 40,
		},
	}


	resp, err := makeRequest("POST", getSitesUrl, token, payload)
	if err != nil{
		fmt.Println("error in making request")

	}
	handleResponse(resp)

}

