package multiservice

import (
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(dsn string) (*UserRepository, error) {
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

}

func getUser(userId int) (User, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/multi_service"
	r, err := NewUserRepository(dsn)	
	if err != nil {
		log.Fatalf("Failed to iniatize user repository: %v", err)
	}


	query := "INSERT INTO users (uuid, name, email, ack) VALUES (?)"

	//execute query 
	res, err := r.db.Exec(query,)
	if err != nil {
		return user, fmt.Errorf("failed to insert user %s into databse: %w", user.Name, err)
	}


}