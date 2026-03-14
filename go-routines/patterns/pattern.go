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

	//creare a chan for communication between the producer and consumer
	ch := make(chan string)

	//wait group
	var wg sync.WaitGroup

	wg.Add(1)
	go simpleProducer(ch, &wg)

	wg.Add(1)
	go simpleConsumer(ch, &wg)

	wg.Wait()

	close(ch)

}

func simpleProducer(ch chan string, wg *sync.WaitGroup) {
	ch <- "Producer sent message"
	wg.Done()
}

func simpleConsumer(ch chan string, wg *sync.WaitGroup) {

	//use select block to have q second timeout
	select {
	case msg := <-ch:
		delay(1)
		fmt.Printf("Consumer Received: %s\n", msg)
	case <-time.After(5 * time.Millisecond):
		fmt.Printf("\nStop listening to channel after %d mili seconds", Timeout)
		wg.Done()
	}

	fmt.Printf("Consumer finishes its job")
	wg.Done()


}


func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}