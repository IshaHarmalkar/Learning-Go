package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string, 3) //buffered chan with capacity 3

	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			orders <- order //only blocks if the buffer is full
			fmt.Println("Placed: ", order)
		}

		close(orders)
	}()

	//Barissa processing orders
	for order := range orders {
		fmt.Printf("Preparing: 5s\n", order)
		time.Sleep(2 * time.Second) //time taken to prepare the order
		fmt.Printf("Served: %s\n", order)
	}

}