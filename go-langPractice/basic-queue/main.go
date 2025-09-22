package main

import (
	"container/list"
	"fmt"
)

func main() {

	//new linked list

	q := list.New()

	//simply append to enque
	q.PushBack(10)
	q.PushBack(20)
	q.PushBack(30)
	
     for e:= q.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	 }

	//deque

	front := q.Front()
	fmt.Println(front.Value)

	q.Remove(front)



}