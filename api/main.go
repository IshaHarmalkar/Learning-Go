package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	apiUrl, token := loadEnv()

	resp, err := makeRequest(apiUrl, token)

	if err != nil {
		log.Fatalf("request failed: %v", err)
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


func makeRequest(apiUrl, token string) (*http.Response, error){
	

	req, err := http.NewRequest("GET", apiUrl, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	//add headers

	req.Header.Set("Authorization", "Bearer " +token)
	req.Header.Set("Accept", "application/json")


	client := &http.Client{Timeout: 10 * time.Second}

	return client.Do(req)
	
}

func handleResponse(resp *http.Response){
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response Body:\n%s\n", string(body))
}

