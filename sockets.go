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

type Connection struct {
	socket    *websocket.Conn
	id        uint
	player_id uint
}

var conns []Connection

func makeConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close() //clean up if we ever exit this function

	var c Connection = Connection{id: connCount, socket: conn}
	conns = append(conns, c)
	connCount = connCount + 1

	//let the client know we are ready to receive messages
	var m = Message{M_CONNECTION_OK, time.Now(), nil}
	err = conn.WriteMessage(websocket.TextMessage, makeMessage(m))

	if err != nil {
		log.Println("write:", err)
		return
	}

	//start listening to messages
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err) //TODO: add player/game id if applicable
			break
		}
		handleMessage(mt, message, err, &c)
	}

}
