package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



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



func (r *UserRepository) Update(user User) error {
	query := "UPDATE students SET name=?, email=? WHERE id=?"
	res, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user ID %d: %w", user.ID, err)
	}
	rowsAffected, _ := res.RowsAffected()
	log.Printf("Updated user ID %d, rows affected: %d", user.ID, rowsAffected)
	return nil
}


func (r *UserRepository) DeleteUser(userID int) error {

	query := "DELETE FROM students WHERE id=?"
	res, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user ID %d: %w", userID, err)
	}

	rowsAffected, _ := res.RowsAffected()
	log.Printf("Deleted user ID %d, rows affected: %d", userID, rowsAffected)
	return nil


}