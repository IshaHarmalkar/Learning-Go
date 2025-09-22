package main

import (
	"fmt"
	"strings"
)

func main() {

	
	
	r := strings.NewReader("hello world")

	


	buf := make([]byte, 5) 


	for{n, err := r.Read(buf)
		if err != nil{
			break
		}

		fmt.Println(string(buf[:n]))
	}

}