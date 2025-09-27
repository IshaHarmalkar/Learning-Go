package main

/* db operation -> write to db, read from db, log kafka msg to db. */

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Id int `json:"id"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type KafkaMessage struct {
	Action string `json:"type"` //create, update, del
	User User `json:"user"`
}




type UserRepository struct {
	db *sql.DB
}


func NewUserRepository(dsn string)(*UserRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connecction: %w", err)
	}
	

	fmt.Println("Inside new user repo fn")

	//pinng the db to ensure the connection is live
	if err:= db.Ping(); err != nil{
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &UserRepository{db: db}, nil

}

//create user
func(r *UserRepository) CreateUser(user User) error {

	fmt.Println("user received in create user: ", user)

	
    ack := false

	uniqId := uuid.Must(uuid.NewRandom()).String()  //uniqId refers to uuid. did not use uuid, as not sure if it would conflcit package name
	fmt.Println("Generate uuid:", uniqId)

	query := "INSERT INTO users (uuid, name, email, ack) VALUES (?, ?, ?, ?)"

	//execute query 
	res, err := r.db.Exec(query,uniqId, user.Name, user.Email, ack)
	if err != nil {
		return fmt.Errorf("failed to insert user %s into databse: %w", user.Name, err)
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch the user id of the user just created: %w", err)
	}

	user.Id = int(userId)

	fmt.Printf("User '%s' successfully created in the db with ID %d and uuid %s. \n", user.Name, userId, user.Uuid)
	log.Printf("Inserted useer with ID: %d", userId)




	r.LogKafkaMsg(user, "create")

    return nil	

}


func(r *UserRepository) LogKafkaMsg(user User, action string) error {

	uniqId := uuid.Must(uuid.NewRandom()).String() 
	fmt.Println("uuid for kafka logs is: ", uniqId)


	jsonData, err := json.Marshal(user)
	if err != nil{
		fmt.Println("error: ", err)
	}

 
	msg := string(jsonData)	
	


    user_id := user.Id


	query := "INSERT INTO kafka_logs (uuid, user_id, msg, action) VALUES (?, ?, ?, ?)"

	//execute query 
	res, err := r.db.Exec(query, uniqId, user_id,  msg, action)
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	if err != nil {
		return fmt.Errorf("failed to %s user %s into databse: %w", action, user.Name, err)
	}

	
	

	fmt.Printf("Logged Kafka msg to db: %v. \n", res)

	km := KafkaMessage{
		Action: action,
		User: user,
	}
	
    //var producerPtr *producer.Producer
	brokers := []string{"localhost:9092"}

	userProducer, err := NewProducer(brokers)

	if err != nil {
		return fmt.Errorf("new producer could not be created in kafka log fn: %v", err)
	}


	//send msg to consumer
	userProducer.SendMsgToConsumer(km)

	defer userProducer.Close()



    return nil	

	

}
