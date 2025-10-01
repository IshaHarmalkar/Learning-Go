package main

import (
	"fmt"
	"time"
)

func main() {

	eventId := 1
	e := Event{
		Id:        eventId,
		Action:    "door_open",
		Timestamp: time.Now(),
		Location: "Bangalore",
	}

	fmt.Println(e)

	u := getUser()
	

	fmt.Println(u)



}