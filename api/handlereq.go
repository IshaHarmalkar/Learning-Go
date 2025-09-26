package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Load global env variables
func loadEnvVars() (string, string) {
	apiURL := os.Getenv("API_URL")
	token := os.Getenv("JWT_TOKEN")

	if apiURL == "" || token == "" {
		log.Fatal("API_URL or JWT_TOKEN not set in environment")
	}

	return apiURL, token
}


// makeRequest handles GET, POST, PATCH, DELETE etc.
func makeRequest(method, url, token string, payload interface{}) (*http.Response, error) {
	var body io.Reader

	


	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewBuffer(data)
	}

	//fmt.Println(payload)
	req, err := http.NewRequest(method, url, body)
	//fmt.Println("Req: ", req)
	//fmt.Println("err: ", err)
	if err != nil {
		fmt.Println("Caught err")

		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 10 * time.Second}

	

	
	//return client.Do(req)
	resp, err:= client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
    
	

	//whatever status code is returned if a token is expired...
	
	if resp.StatusCode == 500{
		log.Fatal("token expired")
	}

	//fmt.Println(resp.Body)
	return resp, err
}

// handleResponse reads and prints the response
func handleResponse(resp *http.Response) {
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {

		
		
		log.Fatalf("Error reading response: %v", err)
	}



	fmt.Printf("Status: %s\n", resp.Status)
	
	fmt.Printf("Response Body:\n%s\n", string(data))
}