package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)


var cacheMutex sync.RWMutex


func getUser(userId string, retry *int, flag *int) (User, int, error) {
	var user User	

	url := "http://localhost:8080/users/" + userId
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		//log.Fatalln("https request failed:", err)
		cacheMutex.Lock()

		*retry++	
		cacheMutex.Unlock()

		*flag = -1

		
		return user, -1, err
	}
	
	

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, 0, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, 0, err
	}

	return user, 1, nil
}

