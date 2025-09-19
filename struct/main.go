package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	//if vars are left undefined, go defines default xero values acc to var type..

	fmt.Println(alex)

	fmt.Printf("%+v", alex)

}