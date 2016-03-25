package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var connCount = 1

var upgrader = websocket.Upgrader{
//default options
}

type Connection struct {
	socket *websocket.Conn
	id     int
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

	brokeGame(&c, connCount)

	for {
		mt, message, err := conn.ReadMessage() //keep listening
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
