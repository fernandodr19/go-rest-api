package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	server := socketio.NewServer(nil)
	
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	
	router.Handle("/socket.io/", server).Methods("GET")
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

