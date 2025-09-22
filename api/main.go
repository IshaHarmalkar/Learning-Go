package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	apiUrl, token := loadEnv()

	//Post
	newUser := User{Name: "Pikachu Pika"}

	resp, err := makePostRequest(apiUrl, token, newUser)

	if err != nil {
		log.Fatalf("Post request failed: %v", err)
	}

	handleResponse(resp)

}

func loadEnv() (string, string) {

	if err := godotenv.Load(); err != nil {
		log.Println("Could not load env file")
	}

	apiUrl := os.Getenv("API_URL")
	token := os.Getenv("JWT_TOKEN")


	if apiUrl == "" || token == ""{
		log.Fatal("Api url or jwt token not set")
	}


	return apiUrl, token

}
