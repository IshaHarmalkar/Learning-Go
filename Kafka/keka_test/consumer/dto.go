package main

import "time"


type Event struct {
	Id        string       `json:"id"`
	EmpCode   int          `json:"empCode"`
	Action    string       `json:"action"`
	Timestamp time.Time    `json:"timestamp"`
	Location  string       `json:"location"`
}

type KafkaMessage struct {
	Event Event `json:"event"`
}