package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	//printGreeting(sb)

}

func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}


// func printGreeting(sp englishBot){
// 	fmt.Println(sb.getGreeting())
// }

func ( englishBot) getGreeting() string {

	//VERY sutom logic for genrating english greeting

	return "Hi There!"
}


func ( spanishBot) getGreeting() string {

	return "Hola!"
}

