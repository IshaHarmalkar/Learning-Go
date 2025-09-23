package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

type UserUpdate struct {
    ID              int       `json:"id"`
    AccessExpiresAt *string   `json:"accessExpiresAt,omitempty"`
    EmployeeCode    *string   `json:"employeeCode,omitempty"`
    Name            *string   `json:"name,omitempty"`
    Email           *string   `json:"email,omitempty"`
    ReportingTo     *string   `json:"reportingTo,omitempty"`
    HomeSiteID      *int      `json:"homeSiteId,omitempty"`
    AdminOfSites    []int     `json:"adminOfSites,omitempty"`
    Roles           []string  `json:"roles,omitempty"`
    Terms           []string  `json:"terms,omitempty"`
    JoiningDate     *string   `json:"joiningDate,omitempty"`
    Attributes      []string  `json:"attributes,omitempty"`
}

type UpdateUsersPayload struct {
    Users []UserUpdate `json:"users"`
}

type DeleteUsersPayload struct {
    UserIDs []int `json:"userIds"`
}



func getSites(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("JWT_TOKEN")

	fmt.Print("Enter the site ID: ")
    reader := bufio.NewReader(os.Stdin)
    siteId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    siteId = strings.TrimSpace(siteId)

	getSitesUrl := os.Getenv("BASE_URL")
	getSitesUrl = getSitesUrl + "/organisationManagement/v2/integrator/organisations/" + siteId+ "/sites"




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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    orgId = strings.TrimSpace(orgId)

	fmt.Print("Enter the site ID: ")
    reader = bufio.NewReader(os.Stdin)
    siteId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    siteId = strings.TrimSpace(siteId)

	url := os.Getenv("BASE_URL")
	url = url + "/organisationManagement/v1/integrator/organisations/" + orgId+ "/sites/" + siteId + "/accessPoints/list"






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

	resp, err := makeRequest("POST", url, token, payload)
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    orgId = strings.TrimSpace(orgId)

	url := os.Getenv("GET_ROLES")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/formData?roles=roles"

	
	


	resp, err := makeRequest("GET", url, token, nil)
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)



	fmt.Println("Enter the user's name: ")
	reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading name: %v", err)
	}

	name := strings.TrimSpace(input)

	
	token := os.Getenv("JWT_TOKEN")

	

	url:= os.Getenv("BASE_URL")

	url = url + "/userManagement/integrator/v1/organisations/" + orgId +"/users/create"

	

	payload := User{
		Name: &name,
	}

	resp, err := makeRequest("POST", url, token, payload)
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)


	url := os.Getenv("BASE_URL")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/users"

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

	resp, err := makeRequest("POST", url, token, payload)
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)

	fmt.Print("Enter the user ID to fetch permission of: ")
    reader = bufio.NewReader(os.Stdin)
    userId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    userId = strings.TrimSpace(userId)

	url := os.Getenv("BASE_URL")

	url = url + "/accessManagementV3/integrator/v1/organisations/" + orgId+"/users/" + userId + "/permissions"
	fmt.Println(url)

	 payload := SitesPayload{
        Sites: []string{}, // An empty slice of strings
    }

	resp, err := makeRequest("POST", url, token, payload)
	if err != nil{
		fmt.Printf("error in making request: %v\n", err)
		return 
	}
	handleResponse(resp)

}




func deactivateUser() {
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)


 
    fmt.Print("Enter the user ID to deactivate: ")


    reader = bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    trimmedInput := strings.TrimSpace(input)
    id, err := strconv.Atoi(trimmedInput)
    if err != nil {
        log.Fatalf("Invalid input: Please enter a valid number. Error: %v", err)
    }

  
    payload := DeactivationPayload{
        Users: []UserDeactivation{
            {
                ID:             id,
                DeactivateUser: true,
            },
        },
    }
    
  
    token := os.Getenv("JWT_TOKEN")
    url := os.Getenv("BASE_URL")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/users/update"

    resp, err := makeRequest("PATCH", url, token, payload)
    if err != nil {
        fmt.Printf("Error in making request: %v\n", err)
        return
    }
    handleResponse(resp)
}


func activateUser() {
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

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)
 
    fmt.Print("Enter the user ID to activate: ")


    reader = bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }

    
    trimmedInput := strings.TrimSpace(input)
    id, err := strconv.Atoi(trimmedInput)
    if err != nil {
        log.Fatalf("Invalid input: Please enter a valid number. Error: %v", err)
    }

  
    payload := DeactivationPayload{
        Users: []UserDeactivation{
            {
                ID:             id,
                DeactivateUser: false,
            },
        },
    }
    
  
    token := os.Getenv("JWT_TOKEN")
    url := os.Getenv("BASE_URL")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/users/update"


    resp, err := makeRequest("PATCH", url, token, payload)
    if err != nil {
        fmt.Printf("Error in making request: %v\n", err)
        return
    }
    handleResponse(resp)
}



func stringPtr(s string) *string {
    return &s
}

func updateUsers() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)
 

    // Prompt for user ID
    fmt.Print("Enter the user ID to update: ")
    reader = bufio.NewReader(os.Stdin)
    inputID, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading ID: %v", err)
    }
    id, err := strconv.Atoi(strings.TrimSpace(inputID))
    if err != nil {
        log.Fatalf("Invalid ID: %v", err)
    }

    // Prompt for user's name
    fmt.Print("Enter the user's name: ")
    inputName, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading name: %v", err)
    }
    name := strings.TrimSpace(inputName)

    // Prompt for user's email
    fmt.Print("Enter the user's email: ")
    inputEmail, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading email: %v", err)
    }
    email := strings.TrimSpace(inputEmail)

    // Create the payload
    payload := UpdateUsersPayload{
        Users: []UserUpdate{
            {
                ID:    id,
                Name:  stringPtr(name),
                Email: stringPtr(email),
                // Other fields are omitted due to omitempty
            },
        },
    }

    // Get auth token and URL from env
    token := os.Getenv("JWT_TOKEN")
    url := os.Getenv("BASE_URL")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/users/update"


    // Make the request
    resp, err := makeRequest("PATCH", url, token, payload)
    if err != nil {
        fmt.Printf("Error in making request: %v\n", err)
        return
    }
    handleResponse(resp)
}



func deleteUser(){

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	fmt.Print("Enter the organisation ID: ")
    reader := bufio.NewReader(os.Stdin)
    orgId, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("Error reading input: %v", err)
    }    
    orgId = strings.TrimSpace(orgId)

	fmt.Println("Enter the user ID to delete: ")
	reader = bufio.NewReader(os.Stdin)

	inputId, err := reader.ReadString('\n')

	

	if err != nil {
		log.Fatalf("Error reading ID: %v", err)
	}

	id, err := strconv.Atoi(strings.TrimSpace(inputId))
	if err != nil {
		log.Fatalf("Invalid ID: Please enter a valid number. Error: %v", err)
	}

	payload := DeleteUsersPayload{
		UserIDs: []int{id},
	}

	token := os.Getenv("JWT_TOKEN")
	url := os.Getenv("BASE_URL")
	url = url + "/userManagement/integrator/v1/organisations/" + orgId + "/users/delete"
	


	resp, err := makeRequest("POST", url, token, payload)

	if err != nil {
		fmt.Printf("Error in making request: %v\n", err)
		return
	}

	handleResponse(resp)

}