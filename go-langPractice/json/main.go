package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type City struct {
	Name       string      `json:"city_name"`
	GDP        int		   `json:"-"`  //does not show field in json output.
	Population int		   `json:"city_population`
}

type User struct {
	Name      string      `json:"name"`
	Age       int 	      `json:"age"`
	City      City        `json:"city"`
	CreatedAt customTime	  `json:"created_at"`
	DeletedAt  customTime  `json:"deleted_at,omitempty"`
}

type customTime struct {
	time.Time
}

const layout = "2006-01-02"



func (c customTime) MarshalJson() ([]byte, error) {

	return []byte(fmt.Sprintf("\"%s\"", c.Format(layout))), nil
}



func main() {

	//t := time.Now()

	u := User{
		Name: "bob",
		Age:  20,
		City: City{
			Name:       "london",
			GDP:        500,
			Population: 8000000},
		CreatedAt: customTime{time.Now()},
		
	}

	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}


