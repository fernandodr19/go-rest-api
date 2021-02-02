package main

import (
	"fmt"
	"time"

	appKafka "./kafka"
)

func main() {
	fmt.Println("Hello Kafka Consumer")
	go appKafka.StartKafka()

	fmt.Println("Kafka consumer has been started")
	time.Sleep(10 * time.Minute)
}
