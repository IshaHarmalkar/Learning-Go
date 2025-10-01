package main

import (
	"time"
)

func main() {

	eventId := 2
	e := Event{
		Id:        eventId,
		Action:    "door_open",
		Timestamp: time.Now(),
		Location: "Bangalore",
	}

	

	u := getUser("1")
	

	/* km := KafkaMessage{
		Event: e,
		User: u,		
	} */

	sendEvent(e, u)



}