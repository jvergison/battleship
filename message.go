package main

import "time"

const (
	M_CONNECTION_OK  string = "Connection ok"
	M_BROKE_NEW_GAME string = "Broker new game"
	M_REJOIN_GAME_OK string = "Rejoin game ok"
)

type Message struct {
	Type      string
	Timestamp time.Time
	Data      map[string]string
}
