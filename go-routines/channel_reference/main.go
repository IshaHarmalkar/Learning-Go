package main

import "fmt"

//sends a msg into existing channel
func sendMessage(ch chan string){
	ch <- "Message sent into  original channel. The chanel that was passed into this fn while calling"

	fmt.Println("len og channel inside fn", len(ch))
}

/* reasign channel to a new channel inside the fn
  The chan from main is passed here too, but instead of appending to it, 
  note how we 
*/

func reassignChannel(ch chan string){
	ch = make(chan string) //new chan, unrealted to original
	ch <- "Message sent into new channel"
	
}

func main(){

	originalChan := make(chan string)

	// sending into original channel 
    go sendMessage(originalChan)

	msg := <-originalChan

	fmt.Println("Received from original channel:", msg)

	

	//wait a second to show program continues
	fmt.Println("Original channel:",len(originalChan))

	


}