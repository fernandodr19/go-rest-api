package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	log = logrus.New()

	// SERVER
	go func() {
		log.WithField("side", "server").Infoln("Go webscokets server")
		http.HandleFunc("/ws", wsEndpoint)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	time.Sleep(2 * time.Second)

	// CLIENT
	log.WithField("side", "client").Infoln("Go webscokets client")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		log.WithField("side", "client").Fatalln("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.WithField("side", "client").WithError(err).Errorln("read error")
				return
			}
			log.WithFields(logrus.Fields{
				"side":    "client",
				"message": string(message),
			}).Infoln("message received")
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.WithField("side", "client").WithError(err).Errorln("failed to write")
				return
			}
		case <-interrupt:
			log.WithField("side", "client").Infoln("interrupted signal received")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.WithField("side", "client").WithError(err).Errorln("write cloese error")
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// SERVER
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Bypass CORS validation
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("side", "server").Fatalln(err)
	}

	log.WithField("side", "server").Infoln("Client Connected")

	reader(ws)
}

func reader(conn *websocket.Conn) {
	go func() {
		for {
			if err := conn.WriteMessage(1, []byte("PING")); err != nil {
				log.WithField("side", "server").Errorln(err)
				return
			}
			time.Sleep(1 * time.Second)
		}

	}()
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.WithField("side", "server").Errorln(err)
			return
		}
		// print out that message for clarity
		log.WithFields(logrus.Fields{
			"side": "server",
			"msg":  string(p),
		})

		if err := conn.WriteMessage(messageType, []byte(`message received 123!`)); err != nil {
			log.WithField("side", "server").Errorln(err)
			return
		}

	}
}
