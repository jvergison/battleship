package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type MessageHandler struct {
	Type        string
	HandlerFunc func(*Connection, Message)
}

var messageHandlers map[string]func(*Connection, Message)

func initMessageHandler() {
	messageHandlers = make(map[string]func(*Connection, Message))
	messageHandlers[M_BROKE_NEW_GAME] = handleBrokeGame
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
	if messageHandlers[m.Type] == nil {
		log.Println("Invalid message received: ", m.Type)
		return
	}
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

func sendMessage(m Message, c *Connection) {
	var err = c.socket.WriteMessage(websocket.TextMessage, makeMessage(m))

	if err != nil {
		log.Println("write:", err)
		return
	}
}

func handleBrokeGame(conn *Connection, message Message) {
	var err error
	var playerId string
	var gameId string
	var action = ""
	var data = make(map[string]string)
	if message.Data != nil {
		err = rejoinGame(conn, message.Data["GameId"], message.Data["PlayerId"])
		action = M_REJOIN_GAME_OK
		if err != nil {
			action = M_FAIL_REJOIN_GAME
			data["Error"] = err.Error()
		}
	} else {
		playerId, gameId, err = brokeNewGame(conn)
		action = M_JOIN_GAME_OK
		if err != nil {
			action = M_FAIL_JOIN_GAME
			data["Error"] = err.Error()
		}
	}

	if err == nil {
		//send success + player id + game id
		data["GameId"] = gameId
		data["PlayerId"] = playerId
	}

	var m = Message{action, time.Now(), data}
	sendMessage(m, conn)

	if action != M_REJOIN_GAME_OK && action != M_FAIL_REJOIN_GAME { //if not rejoining
		checkMatchReady(gameId)
	}

}
