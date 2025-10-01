package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getUser() User {

	resp, err := http.Get("http://localhost:8080/users/1")
	if err != nil {
		log.Fatalln("https request failed:", err)
	}


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	//convert the body to str
	/* sb := string(body)
	log.Printf(sb)
 */

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Converting to struct error: ", err)
	}

	return user


}