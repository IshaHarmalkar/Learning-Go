package multiservice

import "time"

type User struct {
	Id    int    `json:"id"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Event struct {
	Id        int       `json:"id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
}

type KafkaMessage struct {
	Event Event `json:"event"`
	User User   `json:"user"`
}
