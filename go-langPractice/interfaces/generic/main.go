package main

import (
	"fmt"
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




	mysteryBox := interface{}(10) 
	describeValue(mysteryBox)


	retrivedInt, ok := mysteryBox.(string) 

	if ok{
		fmt.Println("Retrived int:", retrivedInt)
	} else{
		fmt.Println("value is not an integer")
	}






}

//can accept anytime
func describeValue(t interface{}){

	fmt.Printf("Type: %T, Value: %v\n", t, t)

}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

