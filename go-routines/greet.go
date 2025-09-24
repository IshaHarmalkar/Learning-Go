package main

import (
	"fmt"
	"time"
)

func greet(ch chan string) {

	time.Sleep(2 * time.Second) //delay for 2 seconds
	ch <- "Hello from go routine!"

}



func main(){
	messageChannel := make(chan string)  //create a channel

	go greet(messageChannel)  //starting a goruitine. Anything with a go keyword in front turns into a go routine

	fmt.Println("Waiting for message...")

    //waiting here for the go rotuines to  finish -> the main go routine stops here until it receives the message
	message := <-messageChannel //recieve message from channel

	fmt.Println(message)
}










/* Pacakge name
Is it a rule of thumb that
all files at the same level in a dir will have the same package.
Also, it is a should or the only way?

*/