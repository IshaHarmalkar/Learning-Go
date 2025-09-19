package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Creaate a ne type of 'deck
//which is a slice of strings

//🌻 deck is an array of string.

type deck []string


func newDeck() deck {

	cards  := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hears", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}


	for _, suit := range cardSuits {
		for _, value := range cardValues{

			cards = append(cards, value + " of " + suit)

		}
	}

	return cards

}




func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}


func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}


func (d deck) toString() string {
	
  return  strings.Join([]string(d), ",")

}


//0666 -> anyone can read and write

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)


}

func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)


	//if no error occured, error = nil.

	if err != nil {
		//option #1 -> log the error and return a call to newDeck()

		//option #2 -> log the error and entirely quite the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}


	s := strings.Split(string(bs), ",")  
	return deck(s)



}
