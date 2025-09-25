package main

import (
	"fmt"
	"time"
)


func placeOrders(orders chan string, numOrders int){
	for i := 1; i <= numOrders; i++{
		order := fmt.Sprintf("Coffee order #%d", i)
		orders <- order //send order to chan
		fmt.Println("Placed:", order)
	}

	close(orders) //signals that no more orders are coming
}

//barrista processing orders
func processOrders(orders chan string){

	//receive from chan until it's closed
	for order := range orders {
		fmt.Printf("Preparing %s\n", order)
		time.Sleep(2 * time.Second) //preparing coffee
		fmt.Printf("Served: %s\n", order)

	}
}

func main() {
	orders := make(chan string) //unbufferd channel

	//start producer goroutine

	go placeOrders(orders, 5)

	processOrders(orders)

	

}


