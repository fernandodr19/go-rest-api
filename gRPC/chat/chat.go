package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// Where is the link between this application and chat.pb.go?

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
