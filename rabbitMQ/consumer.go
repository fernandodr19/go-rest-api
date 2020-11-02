package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer application")

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

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Successfully connected to the RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")

	forever := make(chan bool)
	<-forever
}



