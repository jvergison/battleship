package main

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
