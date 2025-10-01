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
	
	retry := 1
	var u User
	var err error
	flag := -1
	//flag = -1
	//u, flag, err = getUser("1", &retry)

	for retry < 3 && flag == -1 {
		
		fmt.Println("Entered for loop, flag is set to: ", flag)	
		fmt.Println("Entered for loop, retry is set to: ", retry, &retry)		

		u, flag, err := getUser("1", &retry)
		fmt.Printf("Retrying %d with user: %v , flag: %d, and err: %v", retry, u, flag, err)				
		
		if flag == -1 {
		    fmt.Println("sleeping for 1 minute before retrying")
			time.Sleep(30 * time.Second)	

		} else {
			break
		}		   


	}

	if flag == 0 {
		log.Fatalf("Non server error, outside for loop: %v", err)
	} else {
		fmt.Println("User val when server up: ", u)
		sendEvent(e, u)
	}

	

	
}

