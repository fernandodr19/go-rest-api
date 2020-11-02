package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main () {
	fmt.Println("Go Rambbit MQ Publisher Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to Rabbit MQ instance")

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte("Hello world"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully published message do the queue")
}