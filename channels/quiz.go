package main

import "fmt"


func main() {

	c := make(chan string)
	
	for i := 0; i < 4; i ++ {
		go printString("Hello There!", c)   //don't wait for this line to print, continue loop -> punch out multiple go routines
	
	}
	
	
	//infinite loop
	for s := range c {
	
	fmt.Println(s)  // what is s? -> okay, it's like a ele in chan, cool
	
	}
	
}


func printString(s string, c chan string)  {
	fmt.Println(s)
	c <- "Done Printing."

}