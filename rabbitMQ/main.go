package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main () {
	fmt.Println("Go Rambbit MQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to Rabbit MQ instance")
}