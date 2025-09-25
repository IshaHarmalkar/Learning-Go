package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string) //unbufferd channel



	/* since we have a go routine, a new thread is started, main routine 
	by passes this go block
	 */
	go func(){
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #d", i)
			orders <- order //blocks until barista is ready to accept new order
			fmt.Println("Placed:", order)
	}
	close(orders) 


	/* the close tells the barista -> no more orders are comming. 
	 without this the barista's for range orders would block forever.
	*/


	}()


	/* Barista processing orders .
	   Main routine reaches here, and iterates through this loop.


	   in the for loop we are going through the orders channel we made, 
	   that's the link between this for loop and the above go fn.
	
	
	*/
	for order := range orders {
		fmt.Printf("Preparing: %s\n", order)
		time.Sleep(2 * time.Second) //time taken to prepare the order
		fmt.Printf("Served: %s\n", order)



	}

	

}


