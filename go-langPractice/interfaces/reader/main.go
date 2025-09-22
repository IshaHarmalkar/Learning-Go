package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("mango.txt")

	if err != nil {
		fmt.Print("We have some error, exiting next")
		os.Exit(1)
	}

	res, err := io.ReadAll(file)
	
	if err != nil {
			fmt.Print("We have some error, exiting next")
			os.Exit(1)
		}
	fmt.Println(string(res)) // converting to string since the ReadAll returns bytes

}