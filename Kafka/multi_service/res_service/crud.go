package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/mutl_service"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connecction: %w", err)
	}

	fmt.Println("Inside new user repo fn")

	//pinng the db to ensure the connection is live
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("db pointer returnerd")

	return &UserRepository{db: db}, nil

}

func(r *UserRepository) CreatePass(km KafkaMessage) (KafkaMessage, error) {	
    
    fmt.Println("starting create pass: ", km)
	user := km.User
	event := km.Event
	uniqId := uuid.Must(uuid.NewRandom()).String()  //uniqId refers to uuid. did not use uuid, as not sure if it would conflcit package name
	
	query := "INSERT INTO pass (uuid, event_id, pass_action, user_id, name, email, role) VALUES (?, ?, ?, ?, ?, ?, ?)"

	//execute query 
	res, err := r.db.Exec(query,uniqId, event.Id, event.Action,user.Id, user.Name, user.Email, user.Role)
	if err != nil {
		return km, fmt.Errorf("failed to insert user %s into databse: %w", user.Name, err)
	}

	passId, err := res.LastInsertId()
	if err != nil {
		return km, fmt.Errorf("failed to fetch the user id of the user just created: %w", err)
	}

	println("pass id is:%v",passId)

	fmt.Printf("id :%d, pass logged for user %s with event_id %d", passId, user.Name, event.Id)

	
    return km, nil	

}

