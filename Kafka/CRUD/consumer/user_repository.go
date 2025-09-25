package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	query := "INSERT INTO students (name, email) VALUES ( ?, ?)"

	//execute query
	res, err := r.db.Exec(query,user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("fialed to inser user %s into database: %w", user.Name, err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last insert id: %w", err)
	}
	fmt.Printf("User '%s' successfully created in the database with ID %d.\n", user.Name, lastID)
	log.Printf("Inserted user with ID: %d", lastID)

	return nil


}

