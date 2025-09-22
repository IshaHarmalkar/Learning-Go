package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name  string	
}

func makePostRequest(url, token string, payload interface{}) (*http.Response, error) {
	//convert payload to json
	data, err := json.Marshal(payload); 

	if err != nil{
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	//Add headers

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")


	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}



/* func makeGetRequest(url, token string) (*http.Response, error){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
} */


func handleResponse(resp *http.Response){
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response Body: \n%s\n", string(body))

	
}