package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)



type Pagination struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
}

type Filters struct {
    Search   *string  `json:"search,omitempty"`
    UserType []string `json:"userType,omitempty"`
    Terms    []string `json:"terms,omitempty"`
    Sites    []string `json:"sites,omitempty"`
}

type Payload struct {
	Pagination Pagination `json:"pagination"`
	Filters    Filters   `json:"filters,omitempty"`
}

type User struct {
    Name  *string  `json:"name"`
}

type SitesPayload struct {
    Sites []string `json:"sites"`
}

type UserDeactivation struct {
    ID             int  `json:"id"`
    DeactivateUser bool `json:"deactivateUser"`
}

type DeactivationPayload struct {
    Users []UserDeactivation `json:"users"`
}



func getSites(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("JWT_TOKEN")

	getSitesUrl := os.Getenv("POST_SITES")

	payload := Payload{
		Pagination: Pagination{
			Page:    1,
			PerPage: 40,
		},
	}


	resp, err := makeRequest("POST", getSitesUrl, token, payload)
	if err != nil{
		fmt.Println("error in making request")

	}
	handleResponse(resp)

}




func getAccessPoint(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("JWT_TOKEN")

	getSitesUrl := os.Getenv("POST_ACCESS_POINTS")

	var search *string = nil

	payload := Payload{
        Pagination: Pagination{
            Page:    1,
            PerPage: 40,
        },
        Filters: Filters{
            Search: search,
        },
    }

	resp, err := makeRequest("POST", getSitesUrl, token, payload)
	if err != nil{
		fmt.Println("error in making request")

	}
	handleResponse(resp)

}



func getRoles(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("JWT_TOKEN")

	getSitesUrl := os.Getenv("GET_ROLES")

	
	


	resp, err := makeRequest("GET", getSitesUrl, token, nil)
	if err != nil{
		fmt.Println("error in making request")

	}
	handleResponse(resp)

}


func createUser(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := "James Watson"
	token := os.Getenv("JWT_TOKEN")

	Url := os.Getenv("CREATE_USER")

	payload := User{
		Name: &name,
	}

	resp, err := makeRequest("POST", Url, token, payload)
	if err != nil{
		fmt.Printf("error in making request: %v\n", err)
		return // Don't continue if there's an error
	}
	handleResponse(resp)

}





func fetchAllOrgUsers(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	token := os.Getenv("JWT_TOKEN")

	Url := os.Getenv("POST_ALL_USERS")

	payload := Payload{
        Pagination: Pagination{
            Page:    1,
            PerPage: 25,
        },
        Filters: Filters{
            UserType: []string{"active"},
            Terms:    []string{},
            Sites:    []string{},
        },
    }

	resp, err := makeRequest("POST", Url, token, payload)
	if err != nil{
		fmt.Printf("error in making request: %v\n", err)
		return 
	}
	handleResponse(resp)

}





func getUserPermission(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	token := os.Getenv("JWT_TOKEN")

	Url := os.Getenv("POST_USER_PERMISSION")

	 payload := SitesPayload{
        Sites: []string{}, // An empty slice of strings
    }

	resp, err := makeRequest("POST", Url, token, payload)
	if err != nil{
		fmt.Printf("error in making request: %v\n", err)
		return 
	}
	handleResponse(resp)

}


func deactivateUsers(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	type UserDeactivation struct {
        ID             int  `json:"id"`
        DeactivateUser bool `json:"deactivateUser"`
    }

    type DeactivationPayload struct {
        Users []UserDeactivation `json:"users"`
    }
	userID := os.Getenv("USER_ID")
	fmt.Println(userID)

	id, err := strconv.Atoi(userID) 
    if err != nil {
        log.Fatalf("Invalid USER_ID: %v", err)
    }

    payload := DeactivationPayload{
        Users: []UserDeactivation{
            {
                ID: id,
                DeactivateUser: true,
            },
        },
    }
	
	token := os.Getenv("JWT_TOKEN")
	Url := os.Getenv("DEACTIVATE_USER")

	resp, err := makeRequest("PATCH", Url, token, payload)
	if err != nil{
		fmt.Printf("error in making request: %v\n", err)
		return 
	}
	handleResponse(resp)

}