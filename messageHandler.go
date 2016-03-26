package main

import (
	"encoding/json"
	"log"
)

type MessageHandler struct {
	Type        string
	HandlerFunc func(*Connection, Message)
}

var messageHandlers map[string]func(*Connection, Message)

func initMessageHandler() {
	messageHandlers = make(map[string]func(*Connection, Message))
	messageHandlers[M_BROKE_NEW_GAME] = handleBrokeNewGame
}

func handleMessage(messageType int, message []byte, err error, conn *Connection) {
	//convert message to object
	var m Message
	jsonErr := json.Unmarshal(message, &m)
	if jsonErr != nil {
		log.Println("json decode:", jsonErr)
		return
	}
	//call messageHandlers[messagetype]
	messageHandlers[m.Type](conn, m)

}

func makeMessage(m Message) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("json encode:", err)
		return nil
	}

	return b
}

func handleBrokeNewGame(conn *Connection, message Message) {
	brokeGame(conn)

}
