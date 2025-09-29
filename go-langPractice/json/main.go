package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type City struct {
	Name       string
	GDP        int
	Population int
}

type User struct {
	Name      string
	Age       int
	City      City
	CreatedAt time.Time
	DeleteAt  *time.Time
}

func main() {

	u := User{
		Name: "bob",
		Age:  20,
		City: City{
			Name:       "london",
			GDP:        500,
			Population: 8000000},
		CreatedAt: time.Now(),
	}

	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
