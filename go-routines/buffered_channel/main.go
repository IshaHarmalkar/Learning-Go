package main

import (
	"fmt"
	"time"
)

func placeOrders(orders chan string, numOrders int){
	for i := 1; i <= numOrders; i++{
		order := fmt.Sprintf("Coffee order #%d", i)
		orders <- order //blocks only if buffer is full
		fmt.Println("Placed:", order)
	}

	close(orders) //signal no more orders
}


func processOrders(orders chan string){

	//receives unitl chan is closed
	for order := range orders {

		fmt.Printf("Preparin: %s\n", order)
		time.Sleep(2 * time.Second)
		fmt.Printf("Saved: %s\n", order)
		time.Sleep(2 * time.Second) 
		fmt.Printf("Served: %s\n", order)

	}

}

func main() {
	orders := make(chan string, 3) //buffered chan with capacity 3

	//start producer in a go routine
	go placeOrders(orders, 5)

	processOrders(orders)
}


/* 
	The barista will wait if the chan is empty.

	Buffered channels
	-Decouples producers and consumers
	-Burts of traffic
	-Improves throughput
	-Great for worker pools
	-Userful for asynchronous signals or events
*/