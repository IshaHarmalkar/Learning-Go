package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	//duplicates
	e := Event{
		Id:        "b196824e-a6c4-4b03-bb59-dc68a4a0c068",
		Action:    "door_open",
		Timestamp: time.Now(),
		Location:  "Bangalore",
	}

	/*
		eventId := uuid.Must(uuid.NewRandom()).String()
		e := Event{
			Id:         eventId,
			Action:    "door_open",
			Timestamp: time.Now(),
			Location: "Bangalore",
		}

	*/
	retry := 0


	var u User
	var err error
	flag := -1
	//flag = -1
	u, flag, err = getUser("1", &retry, &flag)
	

	for retry < 3 && flag == -1 {

		
		u, flag, err := getUser("1", &retry, &flag)
		fmt.Printf("Retrying %d with user: %v , flag: %d, and err: %v", retry, u, flag, err)				
		cacheMutex.Lock()
		if flag == -1 {
		    fmt.Println("sleeping for 1 minute before retrying")
			time.Sleep(1 * time.Minute)	

		}	
		
	    cacheMutex.Unlock()	


	}

	if flag == 0 {
		log.Fatalf("Non server error: %v", err)
	} else {
		fmt.Println("User val when server up: ", u)
		sendEvent(e, u)
	}

	

	
}

