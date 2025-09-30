package multiservice

import "time"

func main() {

	eventId := 1
	e := Event{
		Id:        eventId,
		Action:    "door_open",
		Timestamp: time.Now(),
		Location: "Bangalore",
	}

}