package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {

	//duplicates
/* 	e := Event{
		Id:        "b196824e-a6c4-4b03-bb59-dc68a4a0c068",
		Action:    "door_open",
		Timestamp: time.Now(),
		Location:  "Bangalore",
	} */

	
		eventId := uuid.Must(uuid.NewRandom()).String()
		e := Event{
			Id:         eventId,
			Action:    "door_open",
			Timestamp: time.Now(),
			Location: "Bangalore",
		}

	
	
	retry := 0
	var u User
	var err error
	flag := -1
	//flag = -1
	//u, flag, err = getUser("1", &retry)

	for retry < 3 && flag == -1 {
		
		/* fmt.Println("Entered for loop, flag is set to: ", flag)	
		fmt.Println("Entered for loop, retry is set to: ", retry, &retry) */		

		u, flag, err = getUser("1", &retry)
		//fmt.Printf("Retrying %d with user: %v , flag: %d, and err: %v", retry, u, flag, err)				
		
		if flag == -1 {
		    fmt.Println("sleeping for 1 minute before retrying")
			time.Sleep(1 * time.Minute)	

		}
	}

	if flag == 0 {
		log.Fatalf("Non server error, outside for loop: %v", err)
	} else if flag == -1 {
		log.Fatalf("Retried 3 times, server still down")

	} else {
		fmt.Println("Server was up, user sent to topic: ", u)
		sendEvent(e, u)
	}

	
}

