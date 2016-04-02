package main

import "time"

const (
	M_CONNECTION_OK    string = "Connection ok"
	M_BROKE_NEW_GAME   string = "Broker new game"
	M_JOIN_GAME_OK     string = "Join game ok"
	M_REJOIN_GAME_OK   string = "Rejoin game ok"
	M_FAIL_REJOIN_GAME string = "Failed rejoin game"
	M_FAIL_JOIN_GAME   string = "Failed join game"
	M_PLAYER_WON       string = "Player won game"
	M_PLAYER_LOST      string = "Player lost game"
)

type Message struct {
	Type      string
	Timestamp time.Time
	Data      map[string]string
}
