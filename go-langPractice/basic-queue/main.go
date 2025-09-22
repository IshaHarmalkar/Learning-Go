package main

import "fmt"

//enque, dequeu, front, rear

type queue []int



func main() {
     
	q := newQueue()

	fmt.Println("Queue Created", q)

	q.enque(5)
	q.dequeu()

	fmt.Print("Element at firt pos in queue is" ,q.first())
	fmt.Println("The last ele in queue is", q.rear())

	fmt.Println("Modified Queue", q)


}

func newQueue() queue {
	q := queue{1, 2, 3, 4}
	return q

}


func (q queue) enque(e int){

	q = append(q, e)
}

func (q queue) dequeu() int {

	ele := q[0]
	q = q[1:]

	return ele

}

func (q queue) first() int {
	return q[0]
}


func (q queue) rear() int {
	return q[len(q) - 1]
}


