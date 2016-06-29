package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var onGoingGames []game

var openGamesMu = &sync.Mutex{}
var openGames []game

var ch = make(chan *connection)

func brokeNewGame(c *connection) (string, string, error) {
	var gameInstance game
	var playerID = randStrings(1)[0]
	c.playerID = playerID
	if len(openGames) == 0 {

		var gameID = randStrings(1)[0]
		gameInstance = game{gameID, c, nil, pStartUp}

		openGamesMu.Lock()
		openGames = append(openGames, gameInstance)
		openGamesMu.Unlock()

		fmt.Printf("player %s started new game %s\n", playerID, gameInstance.id)

	} else {
		gameInstance = joinOpenGame(c)

		onGoingGames = append(onGoingGames, gameInstance)
	}

	c.gameID = gameInstance.id

	return playerID, gameInstance.id, nil

}

func rejoinGame(c *connection, gameID string, playerID string) error {

	var gameInstance, err = findOngoingGameByID(gameID)

	if err == nil {
		if gameInstance.playerOne.playerID == playerID {
			gameInstance.playerOne.socket = c.socket
			ch <- gameInstance.playerOne
		} else if gameInstance.playerTwo.playerID == playerID {
			gameInstance.playerTwo.socket = c.socket
			ch <- gameInstance.playerTwo
		} else {

			return errors.New("Player id expired")
		}

	} else {
		return errors.New("Game does not exist")
	}

	return nil
}

func findOngoingGameByID(id string) (*game, error) {
	for _, gameInstance := range onGoingGames {
		if gameInstance.id == id {
			return &gameInstance, nil
		}
	}

	return nil, errors.New("game not found")

}

func findOpenGameByID(id string) (*game, error) {
	openGamesMu.Lock()
	defer openGamesMu.Unlock()
	for _, gameInstance := range openGames {
		if gameInstance.id == id {
			return &gameInstance, nil
		}
	}

	return nil, errors.New("game not found")

}
func connectionIsInGame(c *connection) bool {
	return c.gameID != ""
}

func joinOpenGame(c *connection) game {

	var gameInstance game

	openGamesMu.Lock()
	gameInstance = openGames[0]
	openGames = openGames[1:] //game will be filled
	openGamesMu.Unlock()

	gameInstance.playerTwo = c
	gameInstance.currentPhase = pPlacement
	fmt.Printf("player %s joined game %s\n", c.playerID, gameInstance.id)

	return gameInstance
}

func onDisconnect(c *connection) {
	//check if in game

	if connectionIsInGame(c) {
		select {
		case m := <-ch:
			//reconnect happened
			fmt.Printf("player %s reconnected\n", m.playerID)
			break
		case <-time.After(3 * time.Minute):
			fmt.Printf("player %s timed out\n", c.playerID)

			//other player wins, throw away game
			var gameInstance, err = findOngoingGameByID(c.gameID)
			if err == nil {
				var otherPlayer *connection
				if gameInstance.playerOne.playerID == c.playerID {
					otherPlayer = gameInstance.playerTwo
				} else {
					otherPlayer = gameInstance.playerOne
				}

				if otherPlayer.playerID != "" {
					//send victory message to player
					var m = message{mPlayerWon, time.Now(), nil}
					sendMessage(m, otherPlayer)
				}
				removePlayer(otherPlayer)

				removeGame(gameInstance)
			}

			removePlayer(c)

			break
		}
	}

}

func removePlayer(c *connection) {
	removefromKnownRandStrings(c.playerID)
	c.playerID = ""
	c.gameID = ""

}

func removeGame(g *game) {
	for i, gameInstance := range onGoingGames {
		if g.id == gameInstance.id {
			onGoingGames = append(onGoingGames[:i], onGoingGames[i+1:]...)
			removefromKnownRandStrings(gameInstance.id)
			return
		}
	}

	openGamesMu.Lock()
	for i, gameInstance := range openGames {
		if g.id == gameInstance.id {
			openGames = append(openGames[:i], openGames[i+1:]...)
			removefromKnownRandStrings(gameInstance.id)
			return
		}
	}
	openGamesMu.Unlock()

}

func checkMatchReady(gameID string) {
	var gameInstance, err = findOpenGameByID(gameID)

	if gameInstance == nil {
		gameInstance, err = findOngoingGameByID(gameID)
	}
	if gameInstance == nil {
		fmt.Printf("game %s not found\n", gameID)
		log.Println("error: ", err)
	}

	if gameInstance != nil {
		if gameInstance.playerOne != nil && gameInstance.playerTwo != nil {
			//start game
			startGame(gameInstance)
		}
	}

}
