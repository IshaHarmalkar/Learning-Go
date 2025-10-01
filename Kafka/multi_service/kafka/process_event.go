package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getUser(userId string, retry *int) (User, int, error) {
	var user User	

	url := "http://localhost:8080/users/" + userId
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		//log.Fatalln("https request failed:", err)
		*retry++		
		return user, -1, err
	}

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

