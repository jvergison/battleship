package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type messageHandler struct {
	mType        string
	handlerFunc func(*connection, message)
}

var messageHandlers map[string]func(*connection, message)

func initMessageHandler() {
	messageHandlers = make(map[string]func(*connection, message))
	messageHandlers[mBrokeNewGame] = handleBrokeGame
}

func handleMessage(messageType int, mess []byte, err error, conn *connection) {
	//convert message to object
	var m message
	jsonErr := json.Unmarshal(mess, &m)
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

func makeMessage(m message) []byte {
	b, err := json.Marshal(m)
	log.Println("sending:",b)
	if err != nil {
		log.Println("json encode:", err)
		return nil
	}

	return b
}

func sendMessage(m message, c *connection) {
	var err = c.socket.WriteMessage(websocket.TextMessage, makeMessage(m))
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func handleBrokeGame(conn *connection, mess message) {

	var err error
	var playerID string
	var gameID string
	var action = ""
	var data = make(map[string]string)
	if mess.Data != nil {
		err = rejoinGame(conn, mess.Data["GameId"], mess.Data["PlayerId"])
		action = mRejoinGameOk
		if err != nil {
			action = mFailRejoinGame
			data["Error"] = err.Error()
		}
	} else {
		playerID, gameID, err = brokeNewGame(conn)
		action = mJoinGameOk
		if err != nil {
			action = mFailJoinGame
			data["Error"] = err.Error()
		}
	}

	if err == nil {
		//send success + player id + game id
		data["GameId"] = gameID
		data["PlayerId"] = playerID
	}

	var m = message{action, time.Now(), data}
	sendMessage(m, conn)

	if action != mRejoinGameOk && action != mFailRejoinGame { //if not rejoining
		checkMatchReady(gameID)
	}

}
