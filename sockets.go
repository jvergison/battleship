package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var connCount uint = 1

var upgrader = websocket.Upgrader{
//default options
}

type connection struct {
	socket       *websocket.Conn
	id           uint
	playerID    string
	gameID      string
	disconnected bool
}

var conns []connection

func makeConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close() //clean up if we ever exit this function

	var c = connection{id: connCount, socket: conn, disconnected: false}
	conns = append(conns, c)
	connCount = connCount + 1

	//let the client know we are ready to receive messages
	var m = message{mConnectionOk, time.Now(), nil}
	sendMessage(m, &c)

	//start listening to messages
	for {
		mt, mess, err := conn.ReadMessage()
		if err != nil {
			c.disconnected = true
			onDisconnect(&c)
			log.Printf("read, player %s : %s", c.playerID, err)
			break
		}
		handleMessage(mt, mess, err, &c)
	}

}
