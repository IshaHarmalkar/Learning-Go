package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	

	//start server & listen
	startServer()

	


	
}

/*

go run main.go producer.go server.go crud.go

*/

func printUser() {

	Jerry := User{
		Id:    1,
		Uuid:  "fjsjkf",
		Name:  "Jerry",
		Email: "jerry.com",
	}

	jsonData, err := json.Marshal(Jerry)
	if err != nil{
		fmt.Println("error: ", err)
	}

    fmt.Println(string(jsonData))

}



