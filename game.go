package main

import (
	"fmt"
	"time"
)

const (
	pStartUp   int = 0 //initial phase value
	pPlacement int = 1 //placement phase
	pPlayer1   int = 2 //player 1 turn
	pPlayer2   int = 3 //player 2 turn
	pFinished  int = 4 //game is over
)

type game struct {
	id           string
	playerOne    *connection
	playerTwo    *connection
	currentPhase int
}

func startGame(game *game) {
	fmt.Printf("game %s starts", game.id)
	game.currentPhase = pPlacement
	var m = message{mPhasePlacement, time.Now(), nil}

	sendMessage(m, game.playerOne)
	sendMessage(m, game.playerTwo)
}
