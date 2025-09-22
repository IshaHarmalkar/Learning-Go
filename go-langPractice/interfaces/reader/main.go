package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	strReader := strings.NewReader("Hellow Go Reader!")

	data, err := io.ReadAll(strReader)
	
	if err != nil {
		print(err)
	}

	fmt.Println(string(data))
}


/* 
strReader := strings.NewReader("Hello Go Reader!")

Creates a *strings.Reader (concrete type).

Implements io.Reader.

data, err := io.ReadAll(strReader)

io.ReadAll is a built-in Go function.

Takes any io.Reader and reads all bytes until EOF.

fmt.Println(string(data))

Prints the string. */