package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//create a program that reads the contents of a text file
	//  then prints its content to the terminal
    fileName := os.Args[1]
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)

	if err!= nil {
		log.Fatal(err)
	} else{
		fmt.Println("File reading possible", f)
		
	}


	io.Copy(os.Stdout, f)





}