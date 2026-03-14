package main

import "fmt"

func main() {

	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	option := 1

	switch option {
	case 1:
		fmt.Println("Play game")
	case 2:
		fmt.Println("Exit Game")
	default:
		fmt.Println("Invalid day")
	}


}