package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var onGoingGames []Game

var openGamesMu = &sync.Mutex{}
var openGames []Game

var ch = make(chan *Connection)

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

	c.game_id = game.id

	return playerId, game.id, nil

}

func rejoinGame(c *Connection, gameId string, playerId string) error {

	var game, err = findGameById(gameId)

	if err == nil {
		if game.PlayerOne.player_id == playerId {
			game.PlayerOne.socket = c.socket
			ch <- game.PlayerOne
		} else if game.PlayerTwo.player_id == playerId {
			game.PlayerTwo.socket = c.socket
			ch <- game.PlayerTwo
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

func connectionIsInGame(c *Connection) bool {
	return c.game_id != ""
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

func onDisconnect(c *Connection) {
	//check if in game

	if connectionIsInGame(c) {
		select {
		case m := <-ch:
			//reconnect happened
			fmt.Printf("player %s reconnected", m.player_id)
			break
		case <-time.After(3 * time.Minute):
			fmt.Printf("player %s timed out", c.player_id)

			//other player wins, throw away game
			var game, err = findGameById(c.game_id)
			if err == nil {
				var otherPlayer *Connection = nil
				if game.PlayerOne.player_id == c.player_id {
					otherPlayer = game.PlayerTwo
				} else {
					otherPlayer = game.PlayerOne
				}

				if otherPlayer.player_id != "" {
					//send victory message to player
					var m = Message{M_PLAYER_WON, time.Now(), nil}
					sendMessage(m, otherPlayer)
				}
				removePlayer(otherPlayer)

				removefromKnownRandStrings(game.id)

				removeGame(game)
			}

			removePlayer(c)

			break
		}
	}

}

func removePlayer(c *Connection) {
	removefromKnownRandStrings(c.player_id)
	c.player_id = ""
	c.game_id = ""

}

func removeGame(g *Game) {
	for i, game := range onGoingGames {
		if g == game {
			onGoingGames = append(onGoingGames[:i], onGoingGames[i+1:]...)
		}
	}

}
