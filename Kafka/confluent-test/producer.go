package main

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	topic           = "message-log"
	broker1Adddress = "localhost:9093"
	broker2Address  = "localhost:9094"
	broker3Address  = "localhost:9095"
)

func produce(ctx context.Context) {

	i := 0

	w := kafka.NewWriter(kafka.WriterConfig {
		Brokers: []string{broker1Adddress, broker2Address, broker3Address},
		Topic: topic,
	})

}