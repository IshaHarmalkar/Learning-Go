package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}




func main() {

	var wg sync.WaitGroup

	//wg := sync.WaitGroup{}
	
	wg.Add(3)


	orders := generateOrders(20)

	// go func() {
	// 	defer wg.Done()
	// 	processOrders(orders)
	// }()


	for i := 0; i <= 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}

	

	


		go reportOrderStatuses(orders)
    

	wg.Wait()

	fmt.Println("All operations completed. Exisitng")


}

func updateOrderStatus(order *Order){

	//update one order at a time
	
		time.Sleep(
			time.Duration(rand.Intn(300)) * 
			time.Millisecond,
		)

		status := []string{
			"Processing", "Shipped", "Delivered",
		}[rand.Intn(3)]

		order.Status = status
		fmt.Printf(
			"Updated order %d status: %s\n",
			order.ID, status,
		)
	

}



func processOrders(orders []*Order){
	for _, order := range orders {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond,)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID: i + 1, Status: "Pending",
		}
	}
	return orders
}

func reportOrderStatuses(orders []*Order){
	for i := 0; i < 5; i ++ {
		fmt.Println("\n-- Order Staatus Report ---")
		for _, order := range orders {
			fmt.Printf(
				"Order %d: %s\n",
				order.ID, order.Status,
			)
		}
		fmt.Printf("-----------------------------\n")
	}
}

