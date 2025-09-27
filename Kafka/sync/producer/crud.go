package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type User struct{
	ID int `json: "id"`
	UUID string `json: "uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type KafkaMessage struct {
	Action string `json:"type"` //create, update, del
	User User `json: "user"`
}




type UserRepository struct {
	db *sql.DB
}


func NewUserRepository(dsn string)(*UserRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connecction: %w", err)
	}

	return &UserRepository{db: db}, nil

}

//create user
func(r *UserRepository) CreateUser(user User) error {

	uniqId := uuid.Must(uuid.NewRandom()).String()  //uniqId refers to uuid. did not use uuid, as not sure if it would conflcit package name
	fmt.Println("Generate uuid:", uniqId)

	query := "INSERT INTO users (uuid, name, email) VALUES (?, ?, ?)"

	//execute query 
	res, err := r.db.Exec(query,uniqId, user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("failed to insert user %s into databse: %w", user.Name, err)
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch the user id of the user just created: %w", err)
	}

	user.ID = int(userId)

	fmt.Printf("User '%s' successfully created in the db with ID %d and uuid %s. \n", user.Name, userId, user.UUID)
	log.Printf("Inserted useer with ID: %d", userId)

	r.LogKafkaMsg(user, "create")

    return nil	

}


func(r *UserRepository) LogKafkaMsg(user User, action string) error {

	uniqId := uuid.Must(uuid.NewRandom()).String() 
	fmt.Println("uuid for kafka logs is: ", uniqId)

	msg, err := json.Marshal(user)

	if err != nil{
		return fmt.Errorf("failed to convert msg to json: %v",  err)
	}


	query := "INSERT INTO kafka_logs (uuid, user_id, msg, action) VALUES (?, ?, ?, ?)"

	//execute query 
	res, err := r.db.Exec(query, uniqId, msg, action)
	if err != nil {
		return fmt.Errorf("failed to insert user %s into databse: %w", user.Name, err)
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch the user id of the user just created: %w", err)
	}

	fmt.Printf("Logged Kafka msg to db. \n",)
	log.Printf("Inserted useer with ID: %d", userId)

    return nil	

}
