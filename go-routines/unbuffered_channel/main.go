package main

import (
	"fmt"
	"time"
)


func placeOrders(orders chan string, numOrders int){
	for i := 1; i <= numOrders; i++{
		order := fmt.Sprintf("Coffee order #%d", i)
		orders <- order //send order to chan, waits till the main rouitne is ready to receive
		fmt.Println("Placed:", order)
	}

	close(orders) //signals that no more orders are coming
}

//barrista processing orders
func processOrders(orders chan string){

	/* receive from chan until it's closed
	ie. waiting on order from orders chan.
	the go routine unblocks over each index?
	
	*/

	/* 
	is this for loop at most on range 1.
	
	*/

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


	//main go routine
	processOrders(orders)

	

}


//Unbuffered channels block the sender until the receiver is ready (and vice versa).