package main

import (
	"time"

	"github.com/google/uuid"
)

func main() {	

	

	
	eventId := uuid.Must(uuid.NewRandom()).String()
	e := Event{
		Id:         eventId,
		EmpCode:   1209,
		Action:    "door_open",
		Timestamp: time.Now(),
		Location: "Bangalore",
	}	

		
	sendEvent(e)
	

	
}
