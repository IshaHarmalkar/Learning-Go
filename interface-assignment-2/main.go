package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//create a program that reads the contents of a text file
	//  then prints its content to the terminal

	f, err := os.OpenFile("myfile.txt", os.O_RDWR|os.O_CREATE, 0644)

	if err!= nil {
		log.Fatal(err)
	} else{
		fmt.Println("File reading possible", f)
	}



}