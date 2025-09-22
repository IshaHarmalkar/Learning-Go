package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}


type Measurable interface{
	Perimerter() float64
}

type Geometry interface{
	Shape
	Measurable

}

type Rectangle struct {
	width, height float64
}


func (r Rectangle) Perimerter() float64 {
	return 2 * (r.width + r.height)
}


func describeShape(g Geometry)  {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimerter())

}



func main() {

	rect := Rectangle{width: 5, height: 4}


	describeShape(rect)




}

type CalculationError struct {
	msg string
}


func (r Rectangle) Area() float64 {
	return r.height * r.width
}


func (ce CalculationError) Error() string{
	return ce.msg
}



func performCaluclation(val float64) (float64, error){
	if  val < 0 {
		return 0 ,CalculationError{
			msg: "Invalid Input",
		}
	}

	return math.Sqrt(val), nil
} 

