package main

import "fmt"

type shape interface{
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t1 := triangle{
		height: 3.0,
		base:   4.0,
	}

	s1 := square{
		sideLength: 4.0,
	}

	printArea(t1)
	printArea(s1)
}



func  printArea(s shape){
	fmt.Println(s.getArea())
}


func (pt triangle) getArea() float64 {
	return 0.5 * pt.base * pt.height
}

func (pt  square) getArea() float64{
	return pt.sideLength * pt.sideLength
}