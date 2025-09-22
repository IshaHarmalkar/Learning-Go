package main

import "fmt"

func main() {

	var queue = make([]int, 0)

	queue = enqueue(queue, 10)
	fmt.Println("After pushing 10 ", queue)


	queue = enqueue(queue, 20)
	fmt.Println("After pushing 20 ", queue)


	queue = enqueue(queue, 30)
	fmt.Println("After pushing 20 ", queue)


	ele, queue := deque(queue)
	fmt.Println("Queue after removing", ele,  " :", queue)

	queue = enqueue(queue, 40)
	fmt.Println("After pushing 40 ", queue)

}

func enqueue(queue []int, element int) []int {

	queue = append(queue, element) //append to q
	fmt.Println("Enqueued:", element)

	return queue

}


func deque(queue []int) (int, []int) {

	element := queue[0]

	
	if len(queue) == 1 {
		var tmp = []int{}
		return element, tmp
	}

	return element, queue[1:]   //slice off the element once it is dequeued

}