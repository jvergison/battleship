package main

import (
	"fmt"
	"time"
)

const (
	P_STARTUP   int = 0
	P_PLACEMENT int = 1
	P_PLAYER1   int = 2
	P_PLAYER2   int = 3
	P_FINISHED  int = 4
)

type Game struct {
	id           string
	PlayerOne    *Connection
	PlayerTwo    *Connection
	currentPhase int
}

func startGame(game *Game) {
	fmt.Printf("game %s starts", game.id)
	game.currentPhase = P_PLACEMENT
	var m = Message{M_PHASE_PLACEMENT, time.Now(), nil}

	sendMessage(m, game.PlayerOne)
	sendMessage(m, game.PlayerTwo)
}
