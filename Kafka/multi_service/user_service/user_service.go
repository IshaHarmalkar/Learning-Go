package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)


type User struct {
	Id    int    `json:"id"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserRepository struct {
	db *sql.DB
}




/* func NewUserRepository() (*UserRepository, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/multi_service"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connecction: %w", err)
	}

	fmt.Println("Inside new user repo fn")

	//pinng the db to ensure the connection is live
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &UserRepository{db: db}, nil

}	  */


func getUser(w http.ResponseWriter,	r *http.Request,) {

	var user User

	//connecting to db
	dsn := "root:@tcp(127.0.0.1:3306)/mutl_service"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("failed to open database connecction: %w", err)
	}

	//pinng the db to ensure the connection is live
	if err := db.Ping(); err != nil {
		log.Panic("failed to ping database: %w", err)
	} 


	//extracting user id from req
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w, err.Error(), http.StatusBadRequest,
		)
		return
	
	}


	//process 

	
	query :="SELECT id, uuid, name, email, role FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	switch err := row.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Role); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(user.Id, user.Uuid)
		default:
			panic(err)

	}

	//fmt.Println(user.Id, user.Uuid)
	w.Header().Set("Content-Type", "application/json")
 	//user.Id = int(id)   
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)

	

}




