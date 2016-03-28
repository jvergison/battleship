package main

import (
	"errors"
	"fmt"
	"sync"
)

var onGoingGames []Game

var openGamesMu = &sync.Mutex{}
var openGames []Game

func brokeNewGame(c *Connection) (string, string, error) {
	var game Game
	var playerId = RandStrings(1)[0]
	c.player_id = playerId
	if len(openGames) == 0 {

		var gameId = RandStrings(1)[0]
		game = Game{gameId, c, nil, P_STARTUP}

		openGamesMu.Lock()
		openGames = append(openGames, game)
		openGamesMu.Unlock()

		fmt.Printf("player %s started new game %s\n", playerId, game.id)

	} else {
		game = joinOpenGame(c)
		onGoingGames = append(onGoingGames, game)
	}

	return playerId, game.id, nil

}

func rejoinGame(c *Connection, gameId string, playerId string) error {

	var game, err = findGameById(gameId)

	if err == nil {
		if game.PlayerOne.player_id == playerId {
			game.PlayerOne.socket = c.socket
		} else if game.PlayerTwo.player_id == playerId {
			game.PlayerTwo.socket = c.socket
		} else {

			return errors.New("Player id expired")
		}

	} else {
		return errors.New("Game does not exist")
	}

	return nil
}

func findGameById(id string) (*Game, error) {
	for _, game := range onGoingGames {
		if game.id == id {
			return &game, nil
		}
	}

	return nil, errors.New("game not found")

}

func joinOpenGame(c *Connection) Game {

	var game Game

	openGamesMu.Lock()
	game = openGames[0]
	openGames = openGames[1:] //game will be filled
	openGamesMu.Unlock()

	game.PlayerTwo = c
	game.currentPhase = P_PLACEMENT
	fmt.Printf("player %s joined game %s\n", c.player_id, game.id)

	return game
}
