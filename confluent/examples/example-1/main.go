package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	singleConsumerProducer()

}

func singleConsumerProducer() {


	//create chan for communication
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go simpleProducer(ch, &wg)

	wg.Add(1)
	go simpleConsumer(ch, &wg)

	wg.Wait()
	close(ch)
}

func simpleProducer(ch chan string, wg *sync.WaitGroup) {
	delay(10)
	ch <- "Producer sent message"
	wg.Done()
}

func simpleConsumer(ch chan string, wg *sync.WaitGroup) {
	
	select {
	case msg := <-ch:
		ConsumerDelay := 1
		delay(ConsumerDelay)
		fmt.Printf("Consumer Received: %s\n", msg)
	
	case <-time.After(10 * time.Millisecond):
		fmt.Printf("\nStop listening to channel after %d milli seconds", 10)
		wg.Done()
	}
	fmt.Printf("Consumer finishes it's job")
	wg.Done()
}


func delay(delay int){
	time.Sleep(time.Duration(delay) * time.Millisecond)
}