package main

import (
	"fmt"
	"time"
)

/* sends a msg after a delay
this is not a 'go routine, it's just a standalone fn. */


func sendMessage(id int, ch chan string){
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Message from goroutine %d", id)
}



func main(){
	messageChan := make(chan string)
	n := 3 //number of go routines to start


	//start multiple go routine

	/* Note how all this channels are communicating with
	 the same channel.
	 Also note, that channel is a reference type.
	*/
	for i := 1; i <= n; i++ {

		/* Go keywords make go routines */
		go sendMessage(i, messageChan)
	}

	fmt.Println("Waiting for messages...")


	//receive messages
	for i := 1; i <= n; i++ {

	//here the main routine is going to have to wait.

		msg := <-messageChan
		fmt.Println(msg)
	}

	fmt.Println("All messages received!")


}