package main

import (
	"fmt"
	"sync"
)

var onGoingGames []Game

var openGamesMu = &sync.Mutex{}
var openGames []Game
var id uint = 1

func brokeGame(c *Connection) bool {
	var game Game
	if len(openGames) == 0 {
		game = Game{id, c, nil, P_STARTUP}

		openGamesMu.Lock()
		openGames = append(openGames, game)
		openGamesMu.Unlock()

		fmt.Printf("player %d started new game %d\n", c.id, game.id)

		id = id + 1

	} else {
		game = joinOpenGame(c)
		onGoingGames = append(onGoingGames, game)
	}

	return true

}

func joinOpenGame(c *Connection) Game {

	var game Game

	openGamesMu.Lock()
	game = openGames[0]
	openGames = openGames[1:] //game will be filled
	openGamesMu.Unlock()

	game.PlayerTwo = c
	game.currentPhase = P_PLACEMENT
	fmt.Printf("player %d joined game %d\n", c.id, game.id)

	return game
}
