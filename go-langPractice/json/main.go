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
	CreatedAt CustomTime	  `json:"created_at"`
	DeletedAt  CustomTime  `json:"deleted_at,omitempty"`
}

type CustomTime  time.Time


const layout = "2006-01-02"



func (c CustomTime) MarshalJson() ([]byte, error) {

	return []byte(c.String()), nil
}

func (c *CustomTime) String() string {
	t := time.Time(*c)
	return fmt.Sprintf("%q", t.Format(layout))
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
		CreatedAt: CustomTime{time.Now()},
		
	}

	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}