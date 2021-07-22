package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"../chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{"x-request-id": "req-123"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	c := chat.NewChatServiceClient(conn)
	// c := chat.ChatServiceClientMock{}

	<-time.After(5 * time.Second)
	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(ctx, &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)

}
