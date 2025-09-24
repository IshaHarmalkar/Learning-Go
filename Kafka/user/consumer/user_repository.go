package main

import (
	"database/sql"
	"fmt"
	"log"
)

//represents a user record in db
/* type User struct {
	ID int
	Name string
	Email string
} */


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(dsn string)(*UserRepository, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	//pinng the db to ensure the connection is live
	if err:= db.Ping(); err != nil{
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &UserRepository{db: db}, nil

}

//creare user inserts a new user record into the users tab;e
func(r *UserRepository) CreateUser(user User) error {
	
	//sql inset query
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"

	//execute query
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("fialed to inser user %s into database: %w", user.Name, err)
	}

	fmt.Printf("User '%s' successfully created in the database.\n", user.Name)
	log.Panicf("Inserted user with ID: %d", user.ID)

	return nil


}

