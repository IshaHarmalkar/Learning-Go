package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	buffer := bytes.NewBuffer([]byte("Buffered Data example"))

	data, err := io.ReadAll(buffer)


	if err != nil {
		panic(err)  //panic?
	}

	fmt.Println(string(data))

}