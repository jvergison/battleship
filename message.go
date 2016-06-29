package main

import "time"

const (
	mConnectionOk    string = "Connection ok"
	mBrokeNewGame   string = "Broker new game"
	mJoinGameOk     string = "Join game ok"
	mRejoinGameOk   string = "Rejoin game ok"
	mFailRejoinGame string = "Failed rejoin game"
	mFailJoinGame   string = "Failed join game"
	mPlayerWon       string = "Player won game"
	mPlayerLost      string = "Player lost game"
	mWaiting          string = "Waiting for player"

	mPhasePlacement string = "Placement phase"
)

type message struct {
	Type      string //must export for json
	Timestamp time.Time //must export for json
	Data      map[string]string //must export for json
}
